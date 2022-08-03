package validations

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPasswordValidatorSuite(t *testing.T) {
	suite.Run(t, new(passwordValidatorSuite))
}

type passwordValidatorSuite struct {
	suite.Suite

	password             string
	passwordConfirmation string

	sut PasswordValidator
}

func (s *passwordValidatorSuite) SetupTest() {
	s.password = "correct horse battery staple"
	s.passwordConfirmation = s.password

	s.sut = NewPasswordValidator()
}

func (s *passwordValidatorSuite) TestValidateShouldReturnErrorIfConfirmationDoesntMatchPassword() {
	// Arrange
	s.passwordConfirmation = s.password + "'"

	// Act
	violations := s.sut.Validate(PasswordValidatorData{
		Password:     s.password,
		Confirmation: s.passwordConfirmation,
	})

	// Assert
	s.Assert().NotEmpty(violations)
	s.Assert().Contains(violations, passwordAndConfirmationMismatch())
}

func (s *passwordValidatorSuite) TestValidateShouldntReturnMismatchWhenConfirmationMatchedPassword() {
	// Act
	violations := s.sut.Validate(PasswordValidatorData{
		Password:     s.password,
		Confirmation: s.passwordConfirmation,
	})

	// Assert
	s.Assert().NotContains(violations, passwordAndConfirmationMismatch())
}
