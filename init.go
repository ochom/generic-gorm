package grm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Platform ...
type Platform string

// Platforms ...
const (
	Postgres Platform = "postgres"
	MySQL    Platform = "mysql"
	Sqlite   Platform = "sqlite"
)

// LogLevels ...
const (
	Silent logger.LogLevel = iota
	Error
	Warn
	Info
)

// Connection ...
type Connection struct {
	SQL *gorm.DB
	Doc *mongo.Database
}

// newConnection ...
func newConnection() *Connection {
	return &Connection{}
}

// withSQL ...
func (c *Connection) withSQL(platform Platform, dsn string, logLevel logger.LogLevel) *Connection {
	var conn *gorm.DB
	var err error

	switch platform {
	case Postgres:
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(0),
		})
	case MySQL:
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel)})
	default:
		conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	}

	if err != nil {
		panic(err)
	}

	c.SQL = conn
	return c
}

// withMongo ...
func (c *Connection) withMongo(dsn, database string) *Connection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	c.Doc = client.Database(database)
	return c
}
