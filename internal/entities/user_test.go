package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u, err := NewUser("test-user@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.ID)
	assert.NotEmpty(t, u.Email)
	assert.NotEmpty(t, u.Password)

	assert.Equal(t, "test-user@gmail.com", u.Email)
	assert.NotEqual(t, "password", u.Password)
	assert.True(t, u.ComparePassword("password"))
	assert.False(t, u.ComparePassword("wrong-password"))

	invalidUser, err := NewUser("invalid-email", "pass")

	assert.NotNil(t, err)
	assert.Nil(t, invalidUser)
}
