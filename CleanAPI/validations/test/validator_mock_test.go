package validationtest

import (
	"fmt"
	"testing"

	v "example.com/api/validations"
	"github.com/stretchr/testify/assert"
)

func ExampleNewValidatorMock() {
	emailValidator := NewValidatorMock(func(email string) []v.Error {
		fmt.Printf("Email is %s\n", email)
		return []v.Error{}
	})

	emailValidator.Validate("some_email@example.com")
	// Output: Email is some_email@example.com
}

func TestEnsureValidatorMockImplementsValidation(t *testing.T) {
	assert.Implements(t, (*v.Validation[any])(nil), new(ValidatorMock[any]))
}
