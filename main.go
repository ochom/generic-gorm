package grm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// this will be initialized by the init function and reused across library
var sql *gorm.DB
var mng *mongo.Client

// Init manually init this package to SQL
func InitSQL(platform Platform, dsn string, logLevel logger.LogLevel) {
	if sql != nil {
		return
	}

	conn := newSQLConnection(platform, dsn, logLevel)
	sql = conn.SQL
}

// InitMongo manually init this package to Mongo
func InitMongo(dsn string) {
	if mng != nil {
		return
	}

	conn := newMongoConnection(dsn)
	mng = conn.Doc
}

// Migrate ...
func Migrate(models ...interface{}) error {
	return sql.AutoMigrate(models...)
}

// SQL get the connection SQL database
func SQL() *gorm.DB {
	return sql
}

// Mongo get the connection Mongo database
func Mongo() *mongo.Client {
	return mng
}
