package init

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/blutspende/bloodlab-common/config"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrFailedToLoadDotEnvFile = errors.New("failed to load .env file")

func LoadDotEnvFile() error {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("%w: %w", ErrFailedToLoadDotEnvFile, err)
		}
	}
	return nil
}

func ReadConfiguration(configuration *config.Configuration) error {
	return config.ReadConfiguration(configuration)
}

func ConfigureLogger(configuration *config.Configuration) {
	consoleWriter := zerolog.NewConsoleWriter()
	consoleWriter.TimeFormat = "2006-01-02T15:04:05Z07:00"
	log.Logger = zerolog.New(consoleWriter).With().Caller().Stack().Timestamp().Logger()
	zerolog.SetGlobalLevel(configuration.ZeroLogLevel)
}

func InitGracefulShutdown() context.Context {
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

func InitStartup(configuration *config.Configuration) (ctx context.Context, err error) {
	err = LoadDotEnvFile()
	if err != nil {
		return nil, err
	}
	err = ReadConfiguration(configuration)
	if err != nil {
		return nil, err
	}
	ConfigureLogger(configuration)
	ctx = InitGracefulShutdown()
	return ctx, nil
}
