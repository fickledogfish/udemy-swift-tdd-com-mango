package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureEmailValidatorImplementsValidation(t *testing.T) {
	assert.Implements(t, (*Validation[string])(nil), new(EmailValidator))
}

func TestValidEmailsShouldReturnNoErrors(t *testing.T) {
	for _, email := range []string{
		"valid_email@example.com",
		"another_valid_email@example.com",
	} {
		t.Logf("Checking %v", email)

		// Arrange
		sut := makeEmailVaidator()

		// Act
		errors := sut.Validate(email)

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
		sut := makeEmailVaidator()

		// Act
		errors := sut.Validate(email)

		// Assert
		assert.NotEmpty(t, errors)
		assert.Contains(t, errors, VErrInvalidEmail{})
	}
}

// Helper functions -----------------------------------------------------------

func makeEmailVaidator() EmailValidator {
	return EmailValidator{}
}
