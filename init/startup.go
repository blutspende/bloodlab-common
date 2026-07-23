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
	configuration         config.Configuration
	buildVersion          string
	UsePostgres           bool
	UseRedis              bool
	UseOtel               bool
	UsePyroscope          bool
	UtcLogging            bool
	startupExtensionFunc  func(config.Configuration) error
	shutdownExtensionFunc func()
}

func Startup(cfg StartupConfig) (ctx context.Context, postgres db.Postgres,
	redisClient *redis.Client, err error) {
	// .env
	err = loadDotEnvFile()
	if err != nil {
		return
	}

	// Configuration
	err = config.ReadConfiguration(cfg.configuration)
	if err != nil {
		return
	}
	commonConfig := cfg.configuration.GetCommonConfig()

	// Logger
	var hook zerolog.Hook
	if cfg.UseOtel {
		hook = correlationIDHook{}
	}
	configureLogger(commonConfig, cfg.UtcLogging, hook)

	// Log startup
	log.Info().Msgf("%s - starting...", commonConfig.ApplicationName)

	// Postgres
	if cfg.UsePostgres {
		postgres = buildPostgres(commonConfig)
	}

	// Redis
	if cfg.UseRedis {
		redisClient = buildRedis(commonConfig)
	}

	// OpenTelemetry
	var tracer *trace.TracerProvider
	var metrics *metric.MeterProvider
	if cfg.UseOtel {
		tracer, metrics = buildOtel(commonConfig, cfg.buildVersion, redisClient)
	}

	// Pyroscope
	var profiler *pyroscope.Profiler
	if cfg.UsePyroscope {
		profiler = buildPyroscope(commonConfig, cfg.buildVersion)
	}

	// Extension function
	err = cfg.startupExtensionFunc(cfg.configuration)
	if err != nil {
		return
	}

	// Graceful shutdown
	ctx = initGracefulShutdown(postgres, redisClient, tracer, metrics, profiler, cfg.shutdownExtensionFunc)

	// Return everything
	return
}
