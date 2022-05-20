package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/api/models"
	v "example.com/api/validations"
	vt "example.com/api/validations/test"
)

func TestEnsureSignUpValidatorImplementsValidation(t *testing.T) {
	assert.Implements(
		t,
		(*v.Validation[models.SignUp])(nil),
		new(SignUpModelValidator),
	)
}

func TestNewSignUpModelValidatorShouldInitializeAllFields(t *testing.T) {
	// Arrange
	sut := makeSignUpModelValidatorSut()

	// Act
	modelValidator := NewSignUpModelValidator(sut.EmailValidator)

	// Assert
	assert.NotNil(t, modelValidator.emailValidator)
}

func TestValidateShouldSendTheCorrectEmailToTheEmailValidator(t *testing.T) {
	// Arrange
	sut := makeSignUpModelValidatorSut()

	expectedEmail := "expected_email@example.com"

	var receivedEmail string
	sut.EmailValidator.ValidateWith = func(email string) []v.Error {
		receivedEmail = email

		return []v.Error{}
	}

	sut.SignUpModel.Email = expectedEmail

	// Act
	sut.SignUpModelValidator.Validate(sut.SignUpModel)

	// Assert
	assert.Equal(t, expectedEmail, receivedEmail)
}

func TestValidateShouldReturnNothingIfEmailValidatorReturnsNothing(t *testing.T) {
	// Arrange
	sut := makeSignUpModelValidatorSut()

	// Act
	errors := sut.SignUpModelValidator.Validate(models.SignUp{})

	// Assert
	assert.Empty(t, errors)
}

func TestValidateShouldReportAllErrorsReportedByTheEmailValidator(t *testing.T) {
	// Arrange
	sut := makeSignUpModelValidatorSut()

	expectedErrors := []v.Error{
		v.VErrInvalidEmail{},
	}

	sut.EmailValidator.ValidateWith = func(string) []v.Error {
		return expectedErrors
	}

	// Act
	errors := sut.SignUpModelValidator.Validate(models.SignUp{})

	// Assert
	assert.NotEmpty(t, errors)
	for _, err := range expectedErrors {
		assert.Contains(t, errors, err)
	}
}

// File SUT -------------------------------------------------------------------

type signUpModelValidatorSut struct {
	SignUpModel          models.SignUp
	EmailValidator       *vt.ValidatorMock[string]
	SignUpModelValidator *SignUpModelValidator
}

func makeSignUpModelValidatorSut() signUpModelValidatorSut {
	emailValidator := vt.NewValidatorMock(func(string) []v.Error {
		return []v.Error{}
	})

	signUpModelValidator := SignUpModelValidator{
		emailValidator: &emailValidator,
	}

	return signUpModelValidatorSut{
		SignUpModel:          models.SignUp{},
		EmailValidator:       &emailValidator,
		SignUpModelValidator: &signUpModelValidator,
	}
}
