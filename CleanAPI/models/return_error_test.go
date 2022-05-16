package models

import (
	"encoding"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureReturnErrorImplementsBinaryMarshaler(t *testing.T) {
	assert.Implements(t, (*encoding.BinaryMarshaler)(nil), new(returnError))
}

func TestNewReturnErrorShouldInitAllFields(t *testing.T) {
	// Arrange
	message := "some error message"

	// Act
	sut := NewReturnError(message)

	// Assert
	assert.Equal(t, message, sut.Error)
}

func TestMarshalBinaryShouldTurnTheStructIntoAJSONString(t *testing.T) {
	// Arrange
	message := "some error message"
	sut := NewReturnError(message)

	// Act
	jsonString, err := sut.MarshalBinary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(
		t,
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		string(jsonString),
	)
}
