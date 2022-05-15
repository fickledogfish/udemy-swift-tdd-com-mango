package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	crypttest "example.com/api/crypt/test"
)

func TestNewUserShouldHashTheUsersPassword(t *testing.T) {
	// Arrange
	name := "some_name"
	email := "some_email@company.site"

	requestModel := SignUp{
		Name: name,
	}

	expectedPasswordHash := "some_hash"

	passwordHasher := crypttest.NewPasswordHasherMock(
		func(string) ([]byte, error) {
			return []byte(expectedPasswordHash), nil
		},
	)

	// Act
	user, err := NewUser(passwordHasher, requestModel)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, string(user.PasswordHash), expectedPasswordHash)
}
