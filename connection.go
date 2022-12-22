package grm

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connection ...
type Connection struct {
	*gorm.DB
}

// Migrate ...
func (c *Connection) Migrate(models ...interface{}) error {
	return c.DB.AutoMigrate(models...)
}

// Platform ...
type Platform string

// Platforms ...
const (
	Postgres Platform = "postgres"
	MySQL    Platform = "mysql"
	Sqlite   Platform = "sqlite"
)

// NewConnection ...
func NewConnection(dbType Platform, dsn string) *Connection {
	var conn *gorm.DB
	var err error
	switch dbType {
	case Postgres:
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case MySQL:
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}

	return &Connection{conn}
}
