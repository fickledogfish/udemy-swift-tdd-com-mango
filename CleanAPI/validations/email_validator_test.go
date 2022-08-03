package validations

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestEmailValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(emailValidatorSuite))
}

type emailValidatorSuite struct {
	suite.Suite

	sut EmailValidator
}

func (s *emailValidatorSuite) SetupTest() {
	s.sut = EmailValidator{}
}

func (s *emailValidatorSuite) TestEnsureEmailValidatorImplementsValidation() {
	s.Assert().Implements((*Validation[string])(nil), s.sut)
}

func (s *emailValidatorSuite) TestValidEmailsShouldReturnNoErrors() {
	for _, email := range []string{
		"valid_email@example.com",
		"another_valid_email@example.com",
	} {
		s.T().Logf("Checking %v", email)

		// Act
		errors := s.sut.Validate(email)

		// Assert
		s.Assert().Empty(errors)
	}
}

func (s *emailValidatorSuite) TestInvalidEmailAddressShouldReturnInvalidEmailError() {
	for _, email := range []string{
		"a",
		"123@",
	} {
		s.T().Logf("Checking %v", email)

		// Act
		errors := s.sut.Validate(email)

		// Assert
		s.Assert().NotEmpty(errors)
		s.Assert().Contains(errors, VErrInvalidEmail{})
	}
}
