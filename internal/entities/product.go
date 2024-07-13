package entities

import (
	"time"

	"github.com/fonsecabc/go-basic-api/pkg/entities"
	"github.com/fonsecabc/go-basic-api/pkg/errors"
)

type Product struct {
	ID        entities.ID `json:"id"`
	Name      string      `json:"name"`
	Price     float32     `json:"price"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewProduct(name string, price float32) (*Product, error) {
	p := &Product{
		ID:        entities.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := p.ValidateProduct(); err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Product) ValidateProduct() error {
	if u.Name == "" {
		return errors.NewMissingParamError("name")
	}

	if u.Price == 0 {
		return errors.NewMissingParamError("price")
	}

	if u.Price < 0 {
		return errors.NewInvalidParamError("price")
	}

	if u.ID.String() == "" {
		return errors.NewMissingParamError("id")
	}

	if _, err := entities.ParseID(u.ID.String()); err != nil {
		return errors.NewInvalidParamError("id")
	}

	return nil
}
