package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/api/models"
	v "example.com/api/validations"
	vt "example.com/api/validations/tests"
)

func TestIsValidShouldSendTheCorrectEmailToTheEmailValidator(t *testing.T) {
	// Arrange
	expectedEmail := "expected_email@example.com"

	var receivedEmail string
	emailValidatorMock := vt.NewEmailValidatorMock(func(email string) []v.Error {
		receivedEmail = email

		return []v.Error{}
	})

	sut := makeSignUpModelValidator(emailValidatorMock)

	// Act
	sut.IsValid(models.SignUpModel{
		Email: expectedEmail,
	})

	// Assert
	assert.Equal(t, expectedEmail, receivedEmail)
}

func TestIsValidShouldReturnNothingIfEmailValidatorReturnsNothing(t *testing.T) {
	// Arrange
	emailValidatorMock := vt.NewEmailValidatorMock(func(string) []v.Error {
		return []v.Error{}
	})
	sut := makeSignUpModelValidator(emailValidatorMock)

	// Act
	errors := sut.IsValid(models.SignUpModel{})

	// Assert
	assert.Empty(t, errors)
}

func TestIsValidShouldReportAllErrorsReportedByTheEmailValidator(t *testing.T) {
	// Arrange
	expectedErrors := []v.Error{
		v.VErrInvalidEmail{},
	}

	emailValidatorMock := vt.NewEmailValidatorMock(func(string) []v.Error {
		return expectedErrors
	})
	sut := makeSignUpModelValidator(emailValidatorMock)

	// Act
	errors := sut.IsValid(models.SignUpModel{})

	// Assert
	assert.NotEmpty(t, errors)
	for _, err := range expectedErrors {
		assert.Contains(t, errors, err)
	}
}

// Helper functions -----------------------------------------------------------

func makeSignUpModelValidator(
	emailValidator v.Validation[string],
) SignUpModelValidator {
	return SignUpModelValidator{
		emailValidator: emailValidator,
	}
}
