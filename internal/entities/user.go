package entities

import (
	"net/mail"

	"github.com/fonsecabc/go-basic-api/pkg/errors"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       value_objects.ID `json:"id"`
	Email    string           `json:"email"`
	Password string           `json:"-"`
}

func NewUser(email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:       value_objects.NewID(),
		Email:    email,
		Password: string(hash),
	}

	if err := u.ValidateUser(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ValidateUser() error {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.NewInvalidParamError("email")
	}

	if len(u.Password) < 6 {
		return errors.NewInvalidParamError("password")
	}

	if u.ID.String() == "" {
		return errors.NewMissingParamError("id")
	}

	if _, err := value_objects.ParseID(u.ID.String()); err != nil {
		return errors.NewInvalidParamError("id")
	}

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
