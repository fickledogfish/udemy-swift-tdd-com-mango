package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBinaryShouldAcceptAValidJSONString(t *testing.T) {
	// Arrange
	expectedSignUp := SignUp{
		Name:                 "Random Person",
		Email:                "person.random@email.com",
		Password:             "123",
		PasswordConfirmation: "321",
	}

	data := []byte(fmt.Sprintf("{"+
		"\"name\": \"%s\","+
		"\"email\": \"%s\","+
		"\"password\": \"%s\","+
		"\"passwordConfirmation\": \"%s\""+
		"}",
		expectedSignUp.Name,
		expectedSignUp.Email,
		expectedSignUp.Password,
		expectedSignUp.PasswordConfirmation,
	))

	var sut SignUp

	// Act
	err := sut.UnmarshalBinary(data)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedSignUp, sut)
}
