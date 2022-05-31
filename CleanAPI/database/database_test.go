package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbt "example.com/api/database/test"
)

func TestEnsureDatabaseImplementsDatabase(t *testing.T) {
	assert.Implements(
		t,
		(*Database[dbt.IdentifiableMock])(nil),
		new(database[dbt.IdentifiableMock]),
	)
}

func TestInsertShouldSendTheIdentifiableToTheChannel(t *testing.T) {
	// Arrange
	sut := makeDatabaseSut()
	identifiableMock := dbt.NewIdentifiableMock("nope")

	// Act
	sut.database.Insert(identifiableMock)
	require.Eventually(
		t,
		func() bool {
			return len(sut.database.eventChan) > 0
		},
		1*time.Second,
		100*time.Millisecond,
	)

	// Assert
	assert.Equal(
		t,
		identifiableMock,
		(<-sut.database.eventChan).data,
	)
}

// File SUT -------------------------------------------------------------------

type databaseSut struct {
	database *database[dbt.IdentifiableMock]

	identifiableMock *dbt.IdentifiableMock
}

func makeDatabaseSut() databaseSut {
	identifiableMock := dbt.NewIdentifiableMock("some_id")

	ch := make(chan event[dbt.IdentifiableMock], 1)

	database := NewDatabaseWithOptions(
		ch,
		func(<-chan event[dbt.IdentifiableMock]) {},
	).(*database[dbt.IdentifiableMock])

	return databaseSut{
		database: database,

		identifiableMock: &identifiableMock,
	}
}
