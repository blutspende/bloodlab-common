package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type DbConnection interface {
	SetSqlConnection(db *sqlx.DB)
	CreateTransactionConnection() (DbConnection, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Rebind(query string) string
	Commit() error
	Rollback() error
	Ping() error
}

type dbConnection struct {
	db              *sqlx.DB
	tx              *sqlx.Tx
	debugLogEnabled bool
}

func NewDbConnection() DbConnection {
	return &dbConnection{}
}

func CreateDbConnector(db *sqlx.DB, enableQueryLogging bool) DbConnection {
	return &dbConnection{
		db:              db,
		debugLogEnabled: enableQueryLogging,
	}
}

func (c *dbConnection) SetSqlConnection(db *sqlx.DB) {
	c.db = db
}

func (c *dbConnection) CreateTransactionConnection() (DbConnection, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		log.Error().Err(err).Msg(MsgBeginTransactionTransactionFailed)
		return nil, ErrBeginTransactionTransactionFailed
	}
	connCopy := *c
	connCopy.tx = tx
	return &connCopy, err
}

func (c *dbConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("Exec query")
	}
	if c.tx != nil {
		return c.tx.Exec(query, args...)
	}
	return c.db.Exec(query, args...)
}

func (c *dbConnection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("ExecContext query")
	}
	if c.tx != nil {
		return c.tx.ExecContext(ctx, query, args...)
	} else {
		return c.db.ExecContext(ctx, query, args...)
	}
}

func (c *dbConnection) Get(dest interface{}, query string, args ...interface{}) error {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("Get query")
	}
	if c.tx != nil {
		return c.tx.Get(dest, query, args...)
	}
	return c.db.Get(dest, query, args...)
}

func (c *dbConnection) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("GetContext query")
	}
	if c.tx != nil {
		return c.tx.GetContext(ctx, dest, query, args...)
	}
	return c.db.GetContext(ctx, dest, query, args...)
}

func (c *dbConnection) NamedExec(query string, arg interface{}) (sql.Result, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("arg", arg).Msg("NamedExec query")
	}
	if c.tx != nil {
		return c.tx.NamedExec(query, arg)
	}
	return c.db.NamedExec(query, arg)
}

func (c *dbConnection) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("arg", arg).Msg("NamedExecContext query")
	}
	if c.tx != nil {
		return c.tx.NamedExecContext(ctx, query, arg)
	}
	return c.db.NamedExecContext(ctx, query, arg)
}

func (c *dbConnection) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("arg", arg).Msg("NamedQuery query")
	}
	if c.tx != nil {
		return c.tx.NamedQuery(query, arg)
	}
	return c.db.NamedQuery(query, arg)
}

func (c *dbConnection) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", arg).Msg("NamedQueryContext query")
	}
	if c.tx != nil {
		return c.tx.NamedQuery(query, arg)
	}
	return c.db.NamedQueryContext(ctx, query, arg)
}

func (c *dbConnection) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("Queryx query")
	}
	if c.tx != nil {
		return c.tx.Queryx(query, args...)
	}
	return c.db.Queryx(query, args...)
}

func (c *dbConnection) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("QueryxContext query")
	}
	if c.tx != nil {
		return c.tx.QueryxContext(ctx, query, args...)
	}
	return c.db.QueryxContext(ctx, query, args...)
}

func (c *dbConnection) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("QueryRowx query")
	}
	if c.tx != nil {
		return c.tx.QueryRowx(query, args...)
	}
	return c.db.QueryRowx(query, args...)
}

func (c *dbConnection) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	if c.debugLogEnabled {
		log.Debug().Str("query", query).Interface("args", args).Msg("QueryRowxContext query")
	}
	if c.tx != nil {
		return c.tx.QueryRowxContext(ctx, query, args...)
	}
	return c.db.QueryRowxContext(ctx, query, args...)
}

func (c *dbConnection) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	if c.tx != nil {
		return c.tx.PrepareNamed(query)
	}
	return c.db.PrepareNamed(query)
}

func (c *dbConnection) Rebind(query string) string {
	if c.tx != nil {
		return c.tx.Rebind(query)
	}
	return c.db.Rebind(query)
}

func (c *dbConnection) Commit() error {
	if c.debugLogEnabled {
		log.Debug().Msg("commit transaction")
	}
	if c.tx == nil {
		// TODO: decide to return error or nil
		return errors.New("invalid transaction, can not perform commit without transaction")
	}
	err := c.tx.Commit()
	c.tx = nil
	if err != nil {
		log.Error().Err(err).Msg(MsgCommitTransactionTransactionFailed)
		return ErrCommitTransactionTransactionFailed
	}
	return nil
}

func (c *dbConnection) Rollback() error {
	if c.debugLogEnabled {
		log.Debug().Msg("rollback transaction")
	}
	if c.tx == nil {
		// TODO: decide to return error or nil
		return errors.New("invalid transaction, can not perform rollback without transaction")
	}
	err := c.tx.Rollback()
	c.tx = nil
	if err != nil {
		log.Error().Err(err).Msg(MsgRollbackTransactionTransactionFailed)
		return ErrRollbackTransactionTransactionFailed
	}
	return nil
}

func (c *dbConnection) Ping() error {
	return c.db.Ping()
}
