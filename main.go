package grm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// this will be initialized by the init function and reused across library
var conn *Connection

// Init manually init this package to SQL
func InitSQL(platform Platform, dsn string, logLevel logger.LogLevel) {
	if conn == nil {
		conn = newConnection()
	}

	if conn.SQL != nil {
		return
	}

	conn = conn.withSQL(platform, dsn, logLevel)
}

// InitMongo manually init this package to Mongo
func InitMongo(dsn, database string) {
	if conn == nil {
		conn = newConnection()
	}

	if conn.Doc != nil {
		return
	}

	conn = conn.withMongo(dsn, database)
}

// Migrate ...
func Migrate(models ...interface{}) error {
	return conn.SQL.AutoMigrate(models...)
}

// SQL get the connection SQL database
func SQL() *gorm.DB {
	return conn.SQL
}

// Mongo get the connection Mongo database
func Mongo() *mongo.Database {
	return conn.Doc
}
