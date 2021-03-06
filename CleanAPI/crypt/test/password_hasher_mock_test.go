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

func TestEnsurePasswordHasherMockImplementsPasswordHasher(t *testing.T) {
	assert.Implements(
		t,
		(*crypt.PasswordHasher)(nil),
		new(PasswordHasherMock),
	)
}
