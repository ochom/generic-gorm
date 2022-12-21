package genericgorm

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// initTestConnection ...
func initTestConnection(t *testing.T) *Connection {
	// sqlite memory
	conn := NewConnection(Sqlite, "file::memory:?cache=shared")
	err := conn.Migrate(&User{})
	if err != nil {
		t.Error(err)
	}

	return conn
}

func getUserRepository(t *testing.T) *Repository[User] {
	conn := initTestConnection(t)
	return NewRepository[User](conn.DB)
}

func TestCreate(t *testing.T) {
	r := getUserRepository(t)
	data := &User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
	}
	err := r.Create(data)
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	r := getUserRepository(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := r.Create(&data)
	require.NoError(t, err)

	user, err := r.GetOne(&User{ID: data.ID})
	require.NoError(t, err)

	user.FirstName = "Ochom Richard"
	err = r.Update(&data)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	r := getUserRepository(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := r.Create(&data)
	require.NoError(t, err)

	err = r.Delete(&data)
	require.NoError(t, err)
}

func TestGetOne(t *testing.T) {
	r := getUserRepository(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := r.Create(&data)
	require.NoError(t, err)

	user, err := r.GetOne(&User{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, user.ID)
}

func TestGetMany(t *testing.T) {
	r := getUserRepository(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := r.Create(&data)
	require.NoError(t, err)

	users, err := r.GetMany(&User{})
	require.NoError(t, err)
	require.NotEmpty(t, users)
}
