package validationtests

import (
	"fmt"

	v "example.com/api/validations"
)

func ExampleNewEmailValidatorMock() {
	emailValidator := NewEmailValidatorMock(func(email string) []v.Error {
		fmt.Printf("Email is %s\n", email)
		return []v.Error{}
	})

	emailValidator.IsValid("some_email@example.com")

	// Output: Email is some_email@example.com
}
