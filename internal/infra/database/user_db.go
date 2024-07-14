package database

import (
	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) LoadById(id value_objects.ID) (*entities.User, error) {
	var user entities.User
	err := u.DB.First(&user, id).Error
	return &user, err
}

func (u *UserDB) LoadByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
