package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/grm"
	"github.com/stretchr/testify/require"
)

func initRepo(t *testing.T) {
	dbType := grm.Sqlite
	dsn := "file::memory:?cache=shared"
	grm.InitSQL(dbType, dsn, grm.Silent)
	if err := grm.Migrate(AllModels...); err != nil {
		t.Error(err)
	}
}

func TestCreate(t *testing.T) {
	initRepo(t)
	data := &User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
	}
	err := grm.Create(data)
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	initRepo(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := grm.Create(&data)
	require.NoError(t, err)

	user, err := grm.GetOne(&User{ID: data.ID})
	require.NoError(t, err)

	user.FirstName = "Ochom Richard"
	err = grm.Update(&data)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	initRepo(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := grm.Create(&data)
	require.NoError(t, err)

	err = grm.Delete(&data)
	require.NoError(t, err)
}

func TestGetOne(t *testing.T) {
	initRepo(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := grm.Create(&data)
	require.NoError(t, err)

	user, err := grm.GetOne(&User{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, user.ID)
}

func TestGetMany(t *testing.T) {
	initRepo(t)
	data := User{
		ID:        uuid.NewString(),
		FirstName: "Ochom",
		LastName:  "Richard",
	}
	err := grm.Create(&data)
	require.NoError(t, err)

	users, err := grm.GetMany(&User{})
	require.NoError(t, err)
	require.NotEmpty(t, users)
}
