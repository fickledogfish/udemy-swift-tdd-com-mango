package database

import "fmt"

type eventType uint

const (
	eventTypeInsert eventType = iota
	eventTypeSelect
)

func (e eventType) String() string {
	switch e {
	case eventTypeInsert:
		return "INSERT"

	case eventTypeSelect:
		return "SELECT"

	default:
		panic(fmt.Sprintf("Unknown event type %d", e))
	}
}

type event[T Identifiable] struct {
	t    eventType
	data T
}

func NewEvent[T Identifiable](t eventType, data T) event[T] {
	return event[T]{
		t:    t,
		data: data,
	}
}

func (e event[T]) String() string {
	return fmt.Sprintf("[%s] %+v", e.t, e.data)
}
