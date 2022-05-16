package validationtest

import (
	"fmt"
	"testing"

	v "example.com/api/validations"
	"github.com/stretchr/testify/assert"
)

func ExampleNewEmailValidatorMock() {
	emailValidator := NewEmailValidatorMock(func(email string) []v.Error {
		fmt.Printf("Email is %s\n", email)
		return []v.Error{}
	})

	emailValidator.Validate("some_email@example.com")

	// Output: Email is some_email@example.com
}

func TestEnsureEmailValidatorMockImplementsValidation(t *testing.T) {
	assert.Implements(t, (*v.Validation[string])(nil), new(emailValidatorMock))
}
