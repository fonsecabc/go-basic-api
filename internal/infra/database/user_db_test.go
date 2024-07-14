package database

import (
	"testing"

	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.User{})

	user, _ := entities.NewUser("test@gmail.com", "password")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	loadedUser, err := userDB.LoadByEmail("test@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, user.ID, loadedUser.ID)
	assert.Equal(t, user.Email, loadedUser.Email)
	assert.Equal(t, user.Password, loadedUser.Password)

	_, err = userDB.LoadByEmail("invalid-email")
	assert.NotNil(t, err)
}

func TestUserDB_LoadById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.User{})

	user, _ := entities.NewUser("test@gmail.com", "password")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	loadedUser, err := userDB.LoadById(user.ID)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, loadedUser.ID)
	assert.Equal(t, user.Email, loadedUser.Email)
	assert.Equal(t, user.Password, loadedUser.Password)
}
