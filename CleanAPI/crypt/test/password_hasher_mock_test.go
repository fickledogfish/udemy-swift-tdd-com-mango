package crypttest

import (
	"fmt"
	"testing"

	"example.com/api/crypt"
	"github.com/stretchr/testify/assert"
)

func ExampleNewPasswordHasherMock() {
	passwordHasher := NewPasswordHasherMock(
		func(password string) ([]byte, error) {
			fmt.Printf("Password is: %s\n", password)
			return []byte{}, nil
		},
	)

	passwordHasher.Hash("some password")
	// Output: Password is: some password
}

func TestEnsurePasswordHasherMockImplementsIPasswordHasher(t *testing.T) {
	assert.Implements(
		t,
		(*crypt.IPasswordHasher)(nil),
		new(passwordHasherMock),
	)
}
