package grm

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection ...
type Connection struct {
	*gorm.DB
}

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

// newConnection ...
func newConnection(dbType Platform, dsn string, logLevel logger.LogLevel) *Connection {
	var conn *gorm.DB
	var err error

	switch dbType {
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

	return &Connection{conn}
}
