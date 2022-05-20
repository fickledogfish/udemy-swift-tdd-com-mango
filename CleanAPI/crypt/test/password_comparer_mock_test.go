package crypttest

import (
	"fmt"
	"testing"

	"example.com/api/crypt"
	"github.com/stretchr/testify/assert"
)

func ExampleNewPasswordComparerMock() {
	passwordComparer := NewPasswordComparerMock(func(hash []byte) bool {
		fmt.Printf("Hash is: %s\n", string(hash))
		return false
	})

	passwordComparer.MatchesHash([]byte("some hash"))
	// Output: Hash is: some hash
}

func TestEnsurePasswordComparerMockImplementsIPasswordComparer(t *testing.T) {
	assert.Implements(
		t,
		(*crypt.IPasswordComparer)(nil),
		new(PasswordComparerMock),
	)
}
