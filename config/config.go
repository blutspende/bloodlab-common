package config

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

var ErrFailedToReadConfiguration = errors.New("failed to read configuration")
var ErrFailedToParseLogLevel = errors.New("failed to parse log level")
var ErrInvalidLogLevel = errors.New("invalid log level")

type Configuration struct {
	ApplicationName string `envconfig:"APPLICATION_NAME" required:"true"`

	PostgresDB struct {
		Host     string `envconfig:"DB_SERVER" required:"true"`
		Port     uint32 `envconfig:"DB_PORT" required:"true"`
		User     string `envconfig:"DB_USER" required:"true"`
		Pass     string `envconfig:"DB_PASS" required:"true"`
		Database string `envconfig:"DB_DATABASE" required:"true"`
		SSLMode  string `envconfig:"DB_SSL_MODE" required:"true"`
	}

	LogLevel     string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	ZeroLogLevel zerolog.Level

	ClientID                        string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret                    string `envconfig:"CLIENT_SECRET" required:"true"`
	ClientCredentialAuthHeaderValue string
}

func ReadConfiguration(configuration *Configuration) error {
	err := envconfig.Process("", configuration)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToReadConfiguration, err)
	}

	zeroLogLevel, err := ParseLogLevel(configuration.LogLevel)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToParseLogLevel, err)
	}
	configuration.ZeroLogLevel = zeroLogLevel

	configuration.ClientCredentialAuthHeaderValue = base64.StdEncoding.EncodeToString([]byte(configuration.ClientID + ":" + configuration.ClientSecret))

	return nil
}

func ParseLogLevel(logLevel string) (zeroLogLevel zerolog.Level, err error) {
	switch strings.ToUpper(logLevel) {
	case "TRACE":
		zeroLogLevel = zerolog.TraceLevel
	case "DEBUG":
		zeroLogLevel = zerolog.DebugLevel
	case "INFO":
		zeroLogLevel = zerolog.InfoLevel
	case "WARN":
		zeroLogLevel = zerolog.WarnLevel
	case "ERROR":
		zeroLogLevel = zerolog.ErrorLevel
	case "FATAL":
		zeroLogLevel = zerolog.FatalLevel
	case "PANIC":
		zeroLogLevel = zerolog.PanicLevel
	case "":
		zeroLogLevel = zerolog.NoLevel
	case "DISABLED":
		zeroLogLevel = zerolog.Disabled
	default:
		return zeroLogLevel, fmt.Errorf("%w: %s", ErrInvalidLogLevel, logLevel)
	}
	return zeroLogLevel, nil
}
