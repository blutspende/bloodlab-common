package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type PgConfig struct {
	ApplicationName              string
	Host                         string
	Port                         uint32
	User                         string
	Pass                         string
	Database                     string
	SSLMode                      string
	MaxOpenConnections           *int
	MaxIdleConnections           *int
	ConnectionMaxLifetimeSeconds *int
	ConnectionMaxIdleTimeSeconds *int
}

type Postgres interface {
	Connect(ctx context.Context) (*sqlx.DB, error)
	GetSqlConnection() (*sqlx.DB, error)
	Close() error
}

type postgres struct {
	config PgConfig
	pgConn *sqlx.DB
}

func NewPostgres(config PgConfig) Postgres {
	return &postgres{
		config: config,
		pgConn: nil,
	}
}

func (p *postgres) Connect(ctx context.Context) (*sqlx.DB, error) {
	if p.pgConn != nil {
		err := p.pgConn.Ping()
		if err == nil {
			log.Warn().Msg("postgres already connected")
			return p.pgConn, nil
		}
	}
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s application_name=%s",
		p.config.Host, p.config.Port, p.config.User,
		p.config.Pass, p.config.Database, p.config.SSLMode, p.config.ApplicationName)
	pgDB, err := sqlx.ConnectContext(ctx, "pgx", url)
	if err != nil {
		log.Error().Err(err).Msg("postgres connection failed")
		return nil, err
	}
	if pgDB == nil {
		log.Error().Msg("postgres connection failed")
		return nil, err
	}
	p.pgConn = pgDB
	if p.config.MaxOpenConnections != nil {
		pgDB.DB.SetMaxOpenConns(*p.config.MaxOpenConnections)
	}
	if p.config.MaxIdleConnections != nil {
		pgDB.DB.SetMaxIdleConns(*p.config.MaxIdleConnections)
	}
	if p.config.ConnectionMaxLifetimeSeconds != nil {
		pgDB.DB.SetConnMaxLifetime(time.Duration(*p.config.ConnectionMaxLifetimeSeconds) * time.Second)
	}
	if p.config.ConnectionMaxIdleTimeSeconds != nil {
		pgDB.DB.SetConnMaxIdleTime(time.Duration(*p.config.ConnectionMaxIdleTimeSeconds) * time.Second)
	}
	log.Info().Msgf("postgres available, connected to %s / %s", p.config.Host, p.config.Database)
	return p.pgConn, nil
}

func (p *postgres) GetSqlConnection() (*sqlx.DB, error) {
	if p.pgConn == nil {
		return nil, fmt.Errorf("postgres connection is not established")
	}
	return p.pgConn, nil
}

func (p *postgres) Close() error {
	if p.pgConn != nil {
		return p.pgConn.Close()
	}
	return nil
}
