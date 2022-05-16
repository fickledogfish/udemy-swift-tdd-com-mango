package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurePasswordComparerImplementsIPasswordComparer(t *testing.T) {
	assert.Implements(
		t,
		(*IPasswordComparer)(nil),
		new(passwordComparer),
	)
}
