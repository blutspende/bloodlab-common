package init

import (
	"context"

	"github.com/blutspende/bloodlab-common/config"
	"github.com/blutspende/bloodlab-common/db"
	"github.com/grafana/pyroscope-go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type StartupConfig struct {
	Configuration         config.Configuration
	BuildVersion          string
	UsePostgres           bool
	UseExtendedPgConfig   bool
	UseRedis              bool
	UseOtel               bool
	UsePyroscope          bool
	UtcLogging            bool
	StartupExtensionFunc  func(config.Configuration) error
	ShutdownExtensionFunc func()
}

func Startup(cfg StartupConfig) (ctx context.Context, dbConn db.DbConnection,
	redisClient *redis.Client, err error) {
	// .env
	err = loadDotEnvFile()
	if err != nil {
		return
	}

	// Configuration
	err = config.ReadConfiguration(cfg.Configuration)
	if err != nil {
		return
	}
	commonConfig := cfg.Configuration.GetCommonConfig()

	// Logger
	var hook zerolog.Hook
	if cfg.UseOtel {
		hook = correlationIDHook{}
	}
	configureLogger(commonConfig, cfg.UtcLogging, hook)

	// Log startup
	log.Info().Msgf("%s - starting...", commonConfig.ApplicationName)

	// Postgres
	var postgres db.Postgres
	if cfg.UsePostgres {
		postgres, dbConn = buildPostgres(commonConfig, cfg.UseExtendedPgConfig)
	}

	// Redis
	if cfg.UseRedis {
		redisClient = buildRedis(commonConfig)
	}

	// OpenTelemetry
	var tracer *trace.TracerProvider
	var metrics *metric.MeterProvider
	if cfg.UseOtel {
		tracer, metrics = buildOtel(commonConfig, cfg.BuildVersion, redisClient)
	}

	// Pyroscope
	var profiler *pyroscope.Profiler
	if cfg.UsePyroscope {
		profiler = buildPyroscope(commonConfig, cfg.BuildVersion)
	}

	// Extension function
	err = cfg.StartupExtensionFunc(cfg.Configuration)
	if err != nil {
		return
	}

	// Graceful shutdown
	ctx = initGracefulShutdown(postgres, redisClient, tracer, metrics, profiler, cfg.ShutdownExtensionFunc)

	// Return everything
	return
}
