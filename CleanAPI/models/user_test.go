package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	crypttest "example.com/api/crypt/test"
)

func TestNewUserShouldHashTheUsersPassword(t *testing.T) {
	// Arrange
	expectedPasswordHash := "some_hash"

	sut := makeUserSut()
	sut.PasswordHasher.CompleteWith = func(string) ([]byte, error) {
		return []byte(expectedPasswordHash), nil
	}

	// Act
	user, err := sut.NewUser()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, sut.Model.Name, user.Name)
	assert.Equal(t, sut.Model.Email, user.Email)
	assert.Equal(t, string(user.PasswordHash), expectedPasswordHash)
}

func TestNewUserShouldReturnErrorIfThePasswordHasherReturnsAnError(t *testing.T) {
	// Arrange
	expectedError := errors.New("C++")

	sut := makeUserSut()
	sut.PasswordHasher.CompleteWith = func(string) ([]byte, error) {
		return []byte{}, expectedError
	}

	// Act
	_, err := sut.NewUser()
	require.Error(t, err)

	// Assert
	assert.ErrorIs(t, err, expectedError)
}

// File SUT -------------------------------------------------------------------

type userSut struct {
	PasswordHasher *crypttest.PasswordHasherMock

	Model *SignUp
}

func makeUserSut() userSut {
	passwordHasher := crypttest.NewPasswordHasherMock(
		func(string) ([]byte, error) {
			return []byte{}, nil
		},
	)

	signUpModel := SignUp{
		Name:                 "some_name",
		Email:                "some_email",
		Password:             "some_password",
		PasswordConfirmation: "some_password_confirmation",
	}

	return userSut{
		PasswordHasher: &passwordHasher,

		Model: &signUpModel,
	}
}

func (sut *userSut) NewUser() (User, error) {
	return NewUser(sut.PasswordHasher, *sut.Model)
}
