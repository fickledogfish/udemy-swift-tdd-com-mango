package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurePasswordComparerImplementsPasswordComparer(t *testing.T) {
	assert.Implements(
		t,
		(*PasswordComparer)(nil),
		new(passwordComparer),
	)
}
