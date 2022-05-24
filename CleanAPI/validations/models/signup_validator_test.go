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
	modelValidator := NewSignUpModelValidator(
		sut.EmailValidator,
		sut.PasswordValidator,
	)

	// Assert
	assert.NotNil(t, modelValidator.emailValidator)
	assert.NotNil(t, modelValidator.passwordValidator)
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

func TestValidateShouldReturnNothingIfAllValidatorsReturnNothing(t *testing.T) {
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

func TestValidateShouldSendTheCorrectPasswordToThePasswordValidator(t *testing.T) {
	// Arrange
	expectedPassword := "some_password"
	expectedConfirmation := "some_password_confirmation"

	sut := makeSignUpModelValidatorSut()
	sut.SignUpModel.Password = expectedPassword
	sut.SignUpModel.PasswordConfirmation = expectedConfirmation

	var receivedPassword string
	var receivedConfirmation string
	sut.PasswordValidator.ValidateWith = func(
		data v.PasswordValidatorData,
	) []v.Error {
		receivedPassword = data.Password
		receivedConfirmation = data.Confirmation

		return []v.Error{}
	}

	// Act
	sut.SignUpModelValidator.Validate(sut.SignUpModel)

	// Assert
	assert.Equal(t, expectedPassword, receivedPassword)
	assert.Equal(t, expectedConfirmation, receivedConfirmation)
}

// File SUT -------------------------------------------------------------------

type signUpModelValidatorSut struct {
	SignUpModel models.SignUp

	EmailValidator    *vt.ValidatorMock[string]
	PasswordValidator *vt.ValidatorMock[v.PasswordValidatorData]

	SignUpModelValidator *SignUpModelValidator
}

func makeSignUpModelValidatorSut() signUpModelValidatorSut {
	emailValidator := vt.NewValidatorMock(func(string) []v.Error {
		return []v.Error{}
	})

	passwordValidator := vt.NewValidatorMock(
		func(v.PasswordValidatorData) []v.Error {
			return []v.Error{}
		},
	)

	signUpModelValidator := NewSignUpModelValidator(
		&emailValidator,
		&passwordValidator,
	)

	return signUpModelValidatorSut{
		SignUpModel: models.SignUp{},

		EmailValidator:    &emailValidator,
		PasswordValidator: &passwordValidator,

		SignUpModelValidator: &signUpModelValidator,
	}
}
