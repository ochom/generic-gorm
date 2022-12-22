package grm

import (
	"log"
	"os"
)

// this will be initialized by the init function and reused across library
var conn *Connection

func getConnectionProps() (Platform, string) {
	dbType := Sqlite
	dsn := "file::memory:?cache=shared"
	db, ok := os.LookupEnv("DATABASE")
	if !ok {
		log.Println("DATABASE not defined, defaulting to SQLITE in memory")
		return dbType, dsn
	}

	dsn, ok = os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("DATABASE_URL not defined, defaulting to SQLITE in memory")
		return dbType, dsn
	}

	return Platform(db), dsn
}

// Init manually init this package
func Init() *Connection {
	platform, dsn := getConnectionProps()
	conn = NewConnection(platform, dsn)
	return conn
}
