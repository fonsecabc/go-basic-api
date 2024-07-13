package entities

import (
	"github.com/fonsecabc/go-basic-api/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entities.ID `json:"id"`
	Email    string      `json:"email"`
	Password string      `json:"-"`
}

func NewUser(email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entities.NewID(),
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
