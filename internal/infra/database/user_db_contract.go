package database

import "github.com/fonsecabc/go-basic-api/internal/entities"

type UserDBContract interface {
	NewUser(user entities.User) error
	LoadUserById(id string) (entities.User, error)
	LoadUserByEmail(email string) (entities.User, error)
}
