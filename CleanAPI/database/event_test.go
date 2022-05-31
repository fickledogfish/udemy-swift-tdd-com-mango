package database

import (
	"fmt"
	"math"
	"testing"

	dbt "example.com/api/database/test"
	"example.com/api/models"

	"github.com/stretchr/testify/assert"
)

func TestEventTypeShouldImplementStringer(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(eventType))
}

func TestEventTypeStringShouldPanicOnUnknownType(t *testing.T) {
	// Assert
	assert.Panics(t, func() {
		// Arrange
		invalidEvent := eventType(math.MaxUint)

		// Act
		_ = invalidEvent.String()
	})
}

func TestEventTypeCasesShouldStringifyToExpectedStrings(t *testing.T) {
	for _, c := range []struct {
		eventType      eventType
		expectedString string
	}{
		{eventTypeInsert, "INSERT"},
		{eventTypeSelect, "SELECT"},
	} {
		assert.Equal(t, c.expectedString, c.eventType.String())
	}
}

func TestEventShouldImplementStringer(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(event[models.User]))
}

func TestEventsShouldStringifyToExpectedStrings(t *testing.T) {
	for _, c := range []struct {
		e              event[dbt.IdentifiableMock]
		expectedString string
	}{{
		NewEvent(eventTypeInsert, dbt.NewIdentifiableMock("some_id")),
		"[INSERT] {Identifier:some_id}",
	}, {
		NewEvent(eventTypeSelect, dbt.NewIdentifiableMock("another_id")),
		"[SELECT] {Identifier:another_id}",
	},
	} {
		assert.Equal(t, c.expectedString, c.e.String())
	}
}
