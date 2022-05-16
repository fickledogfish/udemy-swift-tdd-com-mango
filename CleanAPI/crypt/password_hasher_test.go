package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurePasswordHasherImplementsIPasswordHasher(t *testing.T) {
	assert.Implements(
		t,
		(*IPasswordHasher)(nil),
		new(passwordHasher),
	)
}
