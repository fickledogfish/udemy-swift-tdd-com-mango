package validationtest

import (
	"fmt"
	"testing"

	v "example.com/api/validations"
	"github.com/stretchr/testify/assert"
)

func TestEnsureEmailValidatorMockImplementsValidation(t *testing.T) {
	// Arrange
	var sut interface{} = NewEmailValidatorMock(func(string) []v.Error {
		return []v.Error{}
	})

	// Act
	_, ok := sut.(v.Validation[string])

	// Assert
	assert.True(t, ok)
}

func ExampleNewEmailValidatorMock() {
	emailValidator := NewEmailValidatorMock(func(email string) []v.Error {
		fmt.Printf("Email is %s\n", email)
		return []v.Error{}
	})

	emailValidator.Validate("some_email@example.com")

	// Output: Email is some_email@example.com
}
