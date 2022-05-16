package models

import (
	"encoding"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureSignUpResponseImplementsBinaryMarshaler(t *testing.T) {
	assert.Implements(t, (*encoding.BinaryMarshaler)(nil), new(signUpResponse))
}

func TestNewSignUpResponseShouldInitializeAllFields(t *testing.T) {
	// Arrange
	name := "some_name"
	accessToken := "some_access_token"

	user := User{
		Name:        name,
		AccessToken: accessToken,
	}

	// Act
	sut := NewSignUpResponse(user)

	// Assert
	assert.Equal(t, name, sut.Name)
	assert.Equal(t, accessToken, sut.AccessToken)
}

func TestMarshalBinaryShouldJSONifyTheStruct(t *testing.T) {
	// Arrange
	name := "some_name"
	accessToken := "some_token"

	sut := signUpResponse{
		Name:        name,
		AccessToken: accessToken,
	}

	// Act
	bytes, err := sut.MarshalBinary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf(
		"{"+
			"\"name\":\"%s\","+
			"\"accessToken\":\"%s\""+
			"}",
		name,
		accessToken,
	), string(bytes))
}
