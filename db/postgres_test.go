package db

import (
	"context"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"
)

func TestPostgresConnection(t *testing.T) {
	// Setup embedded Postgres
	embPg := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Port(5551))
	err := embPg.Start()
	assert.Nil(t, err)

	// Create Postgres instance
	dbConfig := PgConfig{
		ApplicationName: "test",
		Host:            "localhost",
		Port:            5551,
		User:            "postgres",
		Pass:            "postgres",
		Database:        "postgres",
		SSLMode:         "disable",
	}
	pg := NewPostgres(dbConfig)

	// Test connection before Connect() is called
	dbConn, err := pg.GetSqlConnection()
	assert.Nil(t, dbConn)
	assert.ErrorIs(t, err, ErrNoPgConnection)

	// Test Connect() method
	dbConn, err = pg.Connect(context.Background())
	assert.NotNil(t, dbConn)
	assert.Nil(t, err)

	// Test GetSqlConnection() after Connect() is called
	dbConn, err = pg.GetSqlConnection()
	assert.NotNil(t, dbConn)
	assert.Nil(t, err)

	// Test Connect() method when already connected
	dbConn, err = pg.Connect(context.Background())
	assert.NotNil(t, dbConn)
	assert.Nil(t, err)

	// Test Close() method
	err = pg.Close()
	assert.Nil(t, err)

	// Test GetSqlConnection() after Close() is called
	dbConn, err = pg.GetSqlConnection()
	assert.Nil(t, dbConn)
	assert.ErrorIs(t, err, ErrNoPgConnection)

	// Stop embedded Postgres
	err = embPg.Stop()
	assert.Nil(t, err)
}
