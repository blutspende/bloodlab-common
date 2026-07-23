package init

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blutspende/bloodlab-common/config"
	"github.com/blutspende/bloodlab-common/db"
	"github.com/grafana/pyroscope-go"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

var ErrFailedToLoadDotEnvFile = errors.New("failed to load .env file")

func loadDotEnvFile() error {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("%w: %w", ErrFailedToLoadDotEnvFile, err)
		}
	}
	return nil
}

func configureLogger(configuration *config.CommonConfiguration, utc bool, hook zerolog.Hook) {
	logLevel := configuration.ZeroLogLevel
	zerolog.SetGlobalLevel(logLevel)

	if utc {
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999999Z"
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().UTC()
		}
	} else {
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05Z07:00"
		zerolog.TimestampFunc = func() time.Time {
			return time.Now()
		}
	}

	consoleWriter := zerolog.NewConsoleWriter()
	consoleLogger := zerolog.New(consoleWriter)
	if hook != nil {
		consoleLogger = consoleLogger.Hook(hook)
	}

	consoleLoggerContext := consoleLogger.With()
	if logLevel <= zerolog.DebugLevel {
		consoleLoggerContext = consoleLoggerContext.Caller()
	}
	consoleLoggerContext = consoleLoggerContext.Caller().Stack().Timestamp()

	log.Logger = consoleLoggerContext.Logger()
	zerolog.DefaultContextLogger = &log.Logger
}

/*
func initGracefulShutdown() context.Context {
	// Init cancelable context for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	// Start wait with graceful shutdown in goroutine
	go func() {
		select {
		case sig := <-sigChan:
			log.Info().Msgf("received termination signal: %+v", sig)
			log.Info().Msg("canceling context")
			cancel() // Cancels the context for all goroutines
		}
	}()
	// Return the cancellable context for use in the application
	return ctx
}
*/

func initGracefulShutdown(postgres db.Postgres, redisClient *redis.Client, tracer *trace.TracerProvider, metrics *metric.MeterProvider, profiler *pyroscope.Profiler, extensionFunc func()) context.Context {
	// Init cancelable context for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM) //nolint

	// TODO: verify if using the same ctx for logs and as the cancellable return is ok
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-sigChan
		log.Info().Ctx(ctx).Msgf("received termination signal: %+v", sig)
		log.Info().Ctx(ctx).Msg("graceful shutdown initiated")

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

		if extensionFunc != nil {
			extensionFunc()
		}

		log.Info().Ctx(ctx).Msg("shutting down")
		os.Exit(0)
	}()

	return ctx
}
