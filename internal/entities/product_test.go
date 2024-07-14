package entities

import (
	"testing"

	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("test", 10.0, value_objects.NewID())

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Name)
	assert.NotEmpty(t, p.Price)
	assert.NotEmpty(t, p.CreatedAt)

	assert.Equal(t, "test", p.Name)
	assert.Equal(t, float32(10.0), p.Price)

	tests := []struct {
		name   string
		price  float32
		userID value_objects.ID
	}{
		{"", 0, value_objects.NewID()},
		{"test", -1, value_objects.NewID()},
		{"test", 0, value_objects.NewID()},
	}

	for _, test := range tests {
		p, err := NewProduct(test.name, test.price, test.userID)
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}
}
