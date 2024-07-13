package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("test", 10.0)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Name)
	assert.NotEmpty(t, p.Price)
	assert.NotEmpty(t, p.CreatedAt)

	assert.Equal(t, "test", p.Name)
	assert.NotEqual(t, 10.0, p.Price)

	// Testing invalid products
	tests := []struct {
		name  string
		price float32
	}{
		{"", 0},
		{"test", -1},
		{"test", 0},
	}

	for _, test := range tests {
		p, err := NewProduct(test.name, test.price)
		assert.NotNil(t, err)
		assert.Nil(t, p)
	}
}
