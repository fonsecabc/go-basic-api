package database

import (
	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
)

type ListFilterBy struct {
	Name string
}

type ProductDBContract interface {
	Create(p entities.Product) error
	LoadById(id value_objects.ID) (entities.Product, error)
	List(userId value_objects.ID, pagination value_objects.PaginationParams[ListFilterBy]) ([]entities.Product, error)
	Update(p entities.Product) error
	Delete(id value_objects.ID) error
}
