package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateShouldReturnErrorIfConfirmationDoesntMatchPassword(t *testing.T) {
	// Arrange
	password := "123"
	passwordConfirmation := "1234"

	passValidator := NewPasswordValidator()

	// Act
	violations := passValidator.Validate(PasswordValidatorData{
		Password:     password,
		Confirmation: passwordConfirmation,
	})

	// Assert
	assert.NotEmpty(t, violations)
	assert.Contains(t, violations, passwordAndConfirmationMismatch())
}

func TestValidateShouldntReturnMismatchWhenConfirmationMatchedPassword(t *testing.T) {
	// Arrange
	password := "123"
	passwordConfirmation := password

	passValidator := NewPasswordValidator()

	// Act
	violations := passValidator.Validate(PasswordValidatorData{
		Password:     password,
		Confirmation: passwordConfirmation,
	})

	// Assert
	assert.NotContains(t, violations, passwordAndConfirmationMismatch())
}
