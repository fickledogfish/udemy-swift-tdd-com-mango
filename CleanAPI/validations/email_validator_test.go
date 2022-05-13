package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotStringsShouldReturnInvalidData(t *testing.T) {
	// Arrange
	sut := makeSut()

	type a struct{}
	notAString := a{}

	// Act
	errors := sut.IsValid(notAString)

	// Assert
	assert.NotEmpty(t, errors)
	assert.Contains(t, errors, VErrInvalidData{})
}

func TestValidEmailsShouldReturnNoErrors(t *testing.T) {
	for _, email := range []string{
		"valid_email@example.com",
		"another_valid_email@example.com",
	} {
		t.Logf("Checking %v", email)

		// Arrange
		sut := makeSut()

		// Act
		errors := sut.IsValid(email)

		// Assert
		assert.Empty(t, errors)
	}
}

func TestInvalidEmailAddressShouldReturnInvalidEmailError(t *testing.T) {
	for _, email := range []string{
		"a",
		"123@",
	} {
		t.Logf("Checking %v", email)

		// Arrange
		sut := makeSut()

		// Act
		errors := sut.IsValid(email)

		// Assert
		assert.NotEmpty(t, errors)
		assert.Contains(t, errors, VErrInvalidEmail{})
	}
}

// Helper functions -----------------------------------------------------------

func makeSut() emailValidator {
	return emailValidator{}
}
