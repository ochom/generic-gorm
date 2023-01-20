package grm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// this will be initialized by the init function and reused across library
var conn *Connection

// Init manually init this package
func Init(platform Platform, dsn string, logLevel logger.LogLevel) *Connection {
	if conn == nil {
		conn = newConnection(platform, dsn, logLevel)
	}
	return conn
}

// Migrate ...
func Migrate(models ...interface{}) error {
	return conn.DB.AutoMigrate(models...)
}

// GetConnection get the connection
func GetConnection() *gorm.DB {
	return conn.DB
}
