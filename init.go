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

// database ...
type database struct {
	sql *gorm.DB
	doc *mongo.Database
}

// newConnection ...
func newConnection() *database {
	return &database{}
}

// withSQL ...
func (c *database) withSQL(platform Platform, dsn string, logLevel logger.LogLevel) *database {
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

	c.sql = conn
	return c
}

// withMongo ...
func (c *database) withMongo(dsn, database string) *database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	c.doc = client.Database(database)
	return c
}
