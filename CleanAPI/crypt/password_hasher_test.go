package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurePasswordHasherImplementsPasswordHasher(t *testing.T) {
	assert.Implements(
		t,
		(*PasswordHasher)(nil),
		new(passwordHasher),
	)
}
