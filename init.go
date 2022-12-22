package grm

// this will be initialized by the init function and reused across library
var conn *Connection

// Init manually init this package
func Init(platform Platform, dsn string) *Connection {
	if conn == nil {
		conn = NewConnection(platform, dsn)
	}
	return conn
}
