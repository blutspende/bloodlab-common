package init

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blutspende/bloodlab-common/config"
	"github.com/blutspende/bloodlab-common/db"
	"github.com/google/uuid"
	"github.com/grafana/pyroscope-go"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.38.0"
)

// TODO: mabye place it somewhere else
const ContextKeyCorrelation = "correlation_id"

type correlationIDHook struct{}

func (h correlationIDHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	if ctx == nil {
		return
	}

	if cid, ok := ctx.Value(ContextKeyCorrelation).(string); ok && cid != "" {
		e.Str("correlationID", cid)
	}
}

func configureLoggerWithMetrics(configuration *config.CommonConfiguration) {
	consoleWriter := zerolog.NewConsoleWriter()
	consoleWriter.TimeFormat = "2006-01-02T15:04:05Z07:00"
	log.Logger = zerolog.New(consoleWriter).Hook(correlationIDHook{}).With().Caller().Stack().Timestamp().Logger()
	zerolog.SetGlobalLevel(configuration.ZeroLogLevel)
}

// Identify the current instance e.g. when running in a cluster the name of the pod
func getInstanceID() string {
	if v := os.Getenv("POD_UID"); v != "" {
		return v // K8s
	}
	if h, _ := os.Hostname(); h != "" {
		return h // otherwise
	}
	return uuid.NewString() //-- and if nothing else is available
}

// Initialize OpenTelemetry tracing and metrics
//
//	traceCollectorEndpoint: host:port of the OTLP trace collector
//	metricsCollectorEndpoint: host:port of the OTLP metrics collector
//
// TODO: verify changes (based on Cerberus implementation)
// - ctx provided externally (same as the cancellable context as used now)
// - instanceID moved from parameter to internally calling getInstanceID instead of at func call
// - configuration parameters directly accessed from config instead of passing them 1by1
func initOpenTelemetry(buildVersion string, configuration *config.CommonConfiguration) (tp *trace.TracerProvider, mp *metric.MeterProvider) {
	ctx := context.Background()

	//-- create resource describing this service
	r, err := resource.New(
		ctx,
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceName(configuration.ApplicationName),
			semconv.ServiceInstanceID(getInstanceID()),
			semconv.ServiceVersion(buildVersion),
		),
	)
	if err != nil {
		// TODO: verify: text changed to "without"
		log.Warn().Err(err).Msg("initialize resource for OpenTelemetry failed: continuing without OpenTelemetry...")
		return nil, nil
	}

	//-- initialize Tracing
	if configuration.OpenTelemetry.TraceCollectorEndpoint != "" {
		traceCollectorEndpoint := configuration.OpenTelemetry.TraceCollectorEndpoint
		exp, err := otlptracegrpc.New(ctx,
			otlptracegrpc.WithInsecure(), // without tls only works for local/secure networks like inside k8s cluster
			otlptracegrpc.WithEndpoint(traceCollectorEndpoint),
			otlptracegrpc.WithTimeout(time.Duration(configuration.OpenTelemetry.TraceCollectorTimeoutSeconds)*time.Second),
		)
		if err != nil {
			log.Warn().Err(err).Msg("initialize OpenTelemetry traces failed. Continuing with OpenTelemetry traces...")
		} else {
			// TODO: these 2 lines come from Lablink, do we need them?
			spanLimits := trace.NewSpanLimits()
			spanLimits.AttributeValueLengthLimit = -1

			tp = trace.NewTracerProvider(
				trace.WithBatcher(exp,
					trace.WithMaxQueueSize(2048),
					trace.WithMaxExportBatchSize(512),
				),
				trace.WithResource(r),
			)
			//-- set global propagator to trace context is required so that the webservice can use it to link the traces to the incoming request
			// this is in the api: e.Use(otelecho.Middleware(serviceName, otelecho.WithTracerProvider(otel.GetTracerProvider()))
			otel.SetTextMapPropagator(propagation.TraceContext{})

			//-- set global tracer. with this tracing already works
			otel.SetTracerProvider(tp)
			log.Info().Msg("OpenTelemetry is enabled, sending traces to " + traceCollectorEndpoint)
		}
	}

	//-- initialize Metrics
	if configuration.OpenTelemetry.MetricsCollectorEndpoint != "" {
		metricsCollectorEndpoint := configuration.OpenTelemetry.MetricsCollectorEndpoint
		expm, err := otlpmetricgrpc.New(
			ctx,
			otlpmetricgrpc.WithEndpoint(metricsCollectorEndpoint),
			otlpmetricgrpc.WithInsecure(),
		)
		if err != nil {
			log.Warn().Err(err).Msg("initialize OpenTelemetry metrics failed. Continuing with OpenTelemetry metrics...")
		} else {
			mp = metric.NewMeterProvider(
				metric.WithResource(r),
				metric.WithReader(
					metric.NewPeriodicReader(expm,
						metric.WithInterval(time.Duration(configuration.OpenTelemetry.MetricsReaderIntervalSeconds)*time.Second),
						metric.WithTimeout(time.Duration(configuration.OpenTelemetry.MetricsReaderTimeoutSeconds)*time.Second),
					),
				),
				// TODO: verify: sdkmetric changed to metric (it was the same package import)
				metric.WithView(
					metric.NewView(
						metric.Instrument{
							Name: "http.server.request.duration",
						},
						metric.Stream{
							Aggregation: metric.AggregationExplicitBucketHistogram{
								Boundaries: []float64{0.05, 0.1, 0.2, 0.3, 0.5, 0.75, 1, 1.5, 2, 3, 5},
							},
						},
					),
				),
				//-- process metrics like CPU, Memory, FD, Threads
				metric.WithView(
					metric.NewView(
						metric.Instrument{
							Name: "process.*",
						},
						metric.Stream{},
					),
				),

				//--  runtime metrics (GC, Goroutines, Heap, Stack)
				metric.WithView(
					metric.NewView(
						metric.Instrument{
							Name: "process.runtime.go.*",
						},
						metric.Stream{},
					),
				),
			)

			otel.SetMeterProvider(mp)
			log.Info().Msg("OpenTelemetry is enabled, sending metrics to " + metricsCollectorEndpoint)
		}

		//-- RAM, gc, goroutines
		if err = runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Duration(configuration.OpenTelemetry.ReadMemStatsIntervalSeconds) * time.Second)); err != nil {
			log.Warn().Err(err).Msg("starting OpenTelemetry MemStats runtime failed")
		}

		//-- Process metrics like CPU, Memory, FD, Threads
		//-- without this process start thing no process metrics are collected
		if err = host.Start(); err != nil {
			log.Warn().Err(err).Msg("starting OpenTelemetry host failed")
		}
	}

	return tp, mp
}

func initGracefulShutdownWithMetrics(postgres db.Postgres, redisClient *redis.Client, tracer *trace.TracerProvider, metrics *metric.MeterProvider, profiler *pyroscope.Profiler) context.Context {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, os.Kill, syscall.SIGTERM) //nolint

	// TODO: verify if using the same ctx for logs and as the cancellable return is ok
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-channel
		log.Info().Ctx(ctx).Msg("graceful shutdown initiated")
		log.Info().Ctx(ctx).Msgf("system call:%+v\n", osCall)

		log.Info().Ctx(ctx).Msg("canceling context")
		cancel()

		if postgres != nil {
			log.Info().Ctx(ctx).Msg("closing DB connection")
			err := postgres.Close()
			if err != nil {
				log.Error().Ctx(ctx).Err(err).Msg("failed to close DB connection")
			}
		}

		if redisClient != nil {
			log.Info().Ctx(ctx).Msg("closing Redis connection")
			if err := redisClient.Close(); err != nil {
				log.Error().Ctx(ctx).Err(err).Msg("failed to close Redis client")
			}
		}

		if profiler != nil {
			log.Info().Ctx(ctx).Msg("stopping Pyroscope profiler")
			profiler.Stop()
		}

		if tracer != nil {
			log.Info().Ctx(ctx).Msg("shutting down OpenTelemetry tracer")
			err := tracer.Shutdown(ctx)
			if err != nil {
				log.Error().Ctx(ctx).Err(err).Msg("failed to shutdown OpenTelemetry tracer")
			}
		}
		if metrics != nil {
			log.Info().Ctx(ctx).Msg("shutting down OpenTelemetry meter")
			err := metrics.Shutdown(ctx)
			if err != nil {
				log.Error().Ctx(ctx).Err(err).Msg("failed to shutdown OpenTelemetry meter")
			}
		}

		log.Info().Ctx(ctx).Msg("shutting down")
		os.Exit(0)
	}()

	return ctx
}

func StartupWithMetrics(configuration config.Configuration, buildVersion string) (ctx context.Context, err error) {
	// .env
	err = loadDotEnvFile()
	if err != nil {
		return nil, err
	}

	// Configuration
	err = config.ReadConfiguration(configuration)

	commonConfig := configuration.GetCommonConfig()

	// Logger
	if err != nil {
		return nil, err
	}
	configureLoggerWithMetrics(commonConfig)

	// Postgres db
	dbConfig := db.PgConfig{
		ApplicationName:              commonConfig.ApplicationName,
		Host:                         commonConfig.PostgresDB.Host,
		Port:                         commonConfig.PostgresDB.Port,
		User:                         commonConfig.PostgresDB.User,
		Pass:                         commonConfig.PostgresDB.Pass,
		Database:                     commonConfig.PostgresDB.Database,
		SSLMode:                      commonConfig.PostgresDB.SSLMode,
		MaxOpenConnections:           new(commonConfig.PostgresDB.MaxOpenConnections),
		MaxIdleConnections:           new(commonConfig.PostgresDB.MaxIdleConnections),
		ConnectionMaxLifetimeSeconds: new(commonConfig.PostgresDB.ConnectionMaxLifetimeSeconds),
		ConnectionMaxIdleTimeSeconds: new(commonConfig.PostgresDB.ConnectionMaxIdleTimeSeconds),
		UseOpenTelemetry:             commonConfig.PostgresDB.UseOpenTelemetry,
	}
	postgres := db.NewPostgres(dbConfig)

	// Redis client
	var redisClient *redis.Client
	if commonConfig.Redis.Enable {
		redisClient = redis.NewClient(&redis.Options{
			Addr:               commonConfig.Redis.Address,
			Protocol:           2,
			Password:           commonConfig.Redis.Password,
			MaxRetries:         commonConfig.Redis.MaxRetries,
			DialerRetries:      commonConfig.Redis.DialerRetries,
			DialerRetryTimeout: time.Duration(commonConfig.Redis.DialerRetryTimeoutMs),
		})
		//defer redisClient.Close() //TODO: verify that this was stupid (graceful shutdown takes care of this)
	}

	// OpenTelemetry tracer and meter
	var tracer *trace.TracerProvider
	var metrics *metric.MeterProvider
	if commonConfig.OpenTelemetry.Enable {
		tracer, metrics = initOpenTelemetry(buildVersion, commonConfig)

		if redisClient != nil {
			if err = redisotel.InstrumentTracing(redisClient); err != nil {
				log.Warn().Err(err).Msg("enable redis opentelemetry tracing failed")
			}
			if err = redisotel.InstrumentMetrics(redisClient); err != nil {
				log.Warn().Err(err).Msg("enable redis opentelemetry metrics failed")
			}
		}
	} else {
		log.Warn().Msg("OpenTelemetry is disabled! This is not recommended for production systems.")
	}

	// Pyroscope Profiling (https://github.com/grafana/pyroscope)
	var profiler *pyroscope.Profiler
	if commonConfig.Pyroscope.Enable {
		profiler, err = pyroscope.Start(pyroscope.Config{
			ApplicationName: commonConfig.ApplicationName,
			ServerAddress:   commonConfig.Pyroscope.Server,
			Tags: map[string]string{
				"service":  commonConfig.ApplicationName,
				"instance": getInstanceID(),
				"version":  buildVersion,
			},
			ProfileTypes: []pyroscope.ProfileType{
				pyroscope.ProfileCPU,
				pyroscope.ProfileAllocObjects,
				pyroscope.ProfileAllocSpace,
				pyroscope.ProfileInuseObjects,
				pyroscope.ProfileInuseSpace,
				//-- mutexes profiling is very expensive, use locally for troubleshooting
				// pyroscope.ProfileMutexCount,
				// pyroscope.ProfileMutexDuration,
			},
		})
		if err != nil {
			log.Warn().Err(err).Msg("Starting Pyroscope profiler failed! Continuing without Pyroscope profiling.")
		}
	}

	// Graceful shutdown
	ctx = initGracefulShutdownWithMetrics(postgres, redisClient, tracer, metrics, profiler)

	return ctx, nil
}
