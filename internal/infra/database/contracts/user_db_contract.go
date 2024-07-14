package database

import (
	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
)

type UserDBContract interface {
	Create(u entities.User) error
	LoadById(id value_objects.ID) (entities.User, error)
	LoadByEmail(email string) (entities.User, error)
}
