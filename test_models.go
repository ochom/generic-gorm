package grm

// User ...
type User struct {
	ID        string
	FirstName string
	LastName  string
}

// AllModels ...
var AllModels = []interface{}{
	&User{},
}
