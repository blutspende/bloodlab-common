package init

import (
	"time"

	"github.com/blutspende/bloodlab-common/config"
	"github.com/blutspende/bloodlab-common/db"
	"github.com/grafana/pyroscope-go"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func buildPostgres(commonConfig *config.CommonConfiguration) db.Postgres {
	dbConfig := db.PgConfig{
		ApplicationName: commonConfig.ApplicationName,
		Host:            commonConfig.PostgresDB.Host,
		Port:            commonConfig.PostgresDB.Port,
		User:            commonConfig.PostgresDB.User,
		Pass:            commonConfig.PostgresDB.Pass,
		Database:        commonConfig.PostgresDB.Database,
		SSLMode:         commonConfig.PostgresDB.SSLMode,
		// TODO: solve the optional addition of these
		MaxOpenConnections:           new(commonConfig.PostgresDB.MaxOpenConnections),
		MaxIdleConnections:           new(commonConfig.PostgresDB.MaxIdleConnections),
		ConnectionMaxLifetimeSeconds: new(commonConfig.PostgresDB.ConnectionMaxLifetimeSeconds),
		ConnectionMaxIdleTimeSeconds: new(commonConfig.PostgresDB.ConnectionMaxIdleTimeSeconds),
		UseOpenTelemetry:             commonConfig.PostgresDB.UseOpenTelemetry,
	}
	return db.NewPostgres(dbConfig)
}

func buildRedis(commonConfig *config.CommonConfiguration) (redisClient *redis.Client) {
	if commonConfig.Redis.Enable {
		redisClient = redis.NewClient(&redis.Options{
			Addr:               commonConfig.Redis.Address,
			Protocol:           2,
			Password:           commonConfig.Redis.Password,
			MaxRetries:         commonConfig.Redis.MaxRetries,
			DialerRetries:      commonConfig.Redis.DialerRetries,
			DialerRetryTimeout: time.Duration(commonConfig.Redis.DialerRetryTimeoutMs),
		})
	} else {
		log.Warn().Msg("Redis is disabled")
	}
	return redisClient
}

func buildOtel(commonConfig *config.CommonConfiguration, buildVersion string, redisClient *redis.Client) (tracer *trace.TracerProvider, metrics *metric.MeterProvider) {
	if commonConfig.OpenTelemetry.Enable {
		tracer, metrics = initOpenTelemetry(buildVersion, commonConfig)
		if redisClient != nil {
			if err := redisotel.InstrumentTracing(redisClient); err != nil {
				log.Warn().Err(err).Msg("enable redis opentelemetry tracing failed")
			}
			if err := redisotel.InstrumentMetrics(redisClient); err != nil {
				log.Warn().Err(err).Msg("enable redis opentelemetry metrics failed")
			}
		}
	} else {
		log.Warn().Msg("OpenTelemetry is disabled - this is not recommended for production systems")
	}
	return tracer, metrics
}

func buildPyroscope(commonConfig *config.CommonConfiguration, buildVersion string) (profiler *pyroscope.Profiler) {
	if commonConfig.Pyroscope.Enable {
		var err error
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
			log.Warn().Err(err).Msg("starting Pyroscope profiler failed - continuing without Pyroscope profiling")
		}
	}
	return profiler
}
