package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	dbt "example.com/api/database/test"
)

func TestEnsureDatabaseImplementsDatabase(t *testing.T) {
	assert.Implements(
		t,
		(*Database[dbt.IdentifiableMock])(nil),
		new(database[dbt.IdentifiableMock]),
	)
}
