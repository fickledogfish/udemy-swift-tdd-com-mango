package models

import (
	"encoding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureSignUpResponseImplementsBinaryMarshaler(t *testing.T) {
	assert.Implements(t, (*encoding.BinaryMarshaler)(nil), new(signUpResponse))
}
