package models

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"example.com/api/models"
	v "example.com/api/validations"
	vt "example.com/api/validations/test"
)

func Test(t *testing.T) {
	suite.Run(t, new(signUpModelValidatorSuite))
}

type signUpModelValidatorSuite struct {
	suite.Suite

	sut SignUpModelValidator

	signUpModel models.SignUp

	emailValidator    *vt.ValidatorMock[string]
	passwordValidator *vt.ValidatorMock[v.PasswordValidatorData]
}

func (s *signUpModelValidatorSuite) SetupTest() {
	s.emailValidator = &vt.ValidatorMock[string]{
		ValidateWith: func(string) []v.Error {
			return []v.Error{}
		},
	}

	s.passwordValidator = &vt.ValidatorMock[v.PasswordValidatorData]{
		ValidateWith: func(v.PasswordValidatorData) []v.Error {
			return []v.Error{}
		},
	}

	s.sut = NewSignUpModelValidator(
		s.emailValidator,
		s.passwordValidator,
	)

	s.signUpModel = models.SignUp{}
}

func (s *signUpModelValidatorSuite) TestEnsureSignUpValidatorImplementsValidation() {
	s.Assert().Implements(
		(*v.Validation[models.SignUp])(nil),
		s.sut,
	)
}

func (s *signUpModelValidatorSuite) TestNewSignUpModelValidatorShouldInitializeAllFields() {
	// Act
	modelValidator := NewSignUpModelValidator(
		s.emailValidator,
		s.passwordValidator,
	)

	// Assert
	s.Assert().NotNil(modelValidator.emailValidator)
	s.Assert().NotNil(modelValidator.passwordValidator)
}

func (s *signUpModelValidatorSuite) TestValidateShouldSendTheCorrectEmailToTheEmailValidator() {
	// Arrange
	expectedEmail := "expected_email@example.com"

	var receivedEmail string
	s.emailValidator.ValidateWith = func(email string) []v.Error {
		receivedEmail = email

		return []v.Error{}
	}

	s.signUpModel.Email = expectedEmail

	// Act
	s.sut.Validate(s.signUpModel)

	// Assert
	s.Assert().Equal(expectedEmail, receivedEmail)
}

func (s *signUpModelValidatorSuite) TestValidateShouldReturnNothingIfAllValidatorsReturnNothing() {
	// Act
	errors := s.sut.Validate(s.signUpModel)

	// Assert
	s.Assert().Empty(errors)
}

func (s *signUpModelValidatorSuite) TestValidateShouldReportAllErrorsReportedByTheEmailValidator() {
	// Arrange
	expectedErrors := []v.Error{
		v.VErrInvalidEmail{},
	}

	s.emailValidator.ValidateWith = func(string) []v.Error {
		return expectedErrors
	}

	// Act
	errors := s.sut.Validate(s.signUpModel)

	// Assert
	s.Assert().NotEmpty(errors)
	for _, err := range expectedErrors {
		s.Assert().Contains(errors, err)
	}
}

func (s *signUpModelValidatorSuite) TestValidateShouldSendTheCorrectPasswordToThePasswordValidator() {
	// Arrange
	expectedPassword := "some_password"
	expectedConfirmation := "some_password_confirmation"

	s.signUpModel.Password = expectedPassword
	s.signUpModel.PasswordConfirmation = expectedConfirmation

	var receivedPassword string
	var receivedConfirmation string
	s.passwordValidator.ValidateWith = func(
		data v.PasswordValidatorData,
	) []v.Error {
		receivedPassword = data.Password
		receivedConfirmation = data.Confirmation

		return []v.Error{}
	}

	// Act
	s.sut.Validate(s.signUpModel)

	// Assert
	s.Assert().Equal(expectedPassword, receivedPassword)
	s.Assert().Equal(expectedConfirmation, receivedConfirmation)
}
