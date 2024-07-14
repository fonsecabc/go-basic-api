package entities

import (
	"time"

	"github.com/fonsecabc/go-basic-api/pkg/errors"
	"github.com/fonsecabc/go-basic-api/pkg/validations"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
)

type Product struct {
	ID        value_objects.ID `json:"id"`
	Name      string           `json:"name"`
	Price     float32          `json:"price"`
	UserID    value_objects.ID `json:"user_id"`
	CreatedAt time.Time        `json:"created_at"`
}

func NewProduct(name string, price float32, userID value_objects.ID) (*Product, error) {
	p := &Product{
		ID:        value_objects.NewID(),
		Name:      name,
		Price:     price,
		UserID:    userID,
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

	if err := validations.ValidateID(u.ID); err != nil {
		return err
	}

	if err := validations.ValidateID(u.UserID); err != nil {
		return err
	}

	return nil
}
