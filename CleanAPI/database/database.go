package database

import "fmt"

type Database[T Identifiable] interface {
	Inserter[T]
}

type database[T Identifiable] struct {
	Data []T

	eventChan chan event[T]
}

func NewDatabase[T Identifiable]() Database[T] {
	ch := make(chan event[T], 10)

	eventObserver := func(ch <-chan event[T]) {
		for event := range ch {
			fmt.Println(event)
		}
	}

	return NewDatabaseWithOptions(ch, eventObserver)
}

func NewDatabaseWithOptions[T Identifiable](
	eventChan chan event[T],
	eventObserver func(<-chan event[T]),
) Database[T] {
	go eventObserver(eventChan)

	return &database[T]{
		Data: []T{},

		eventChan: eventChan,
	}
}

// Implementing Inserter ------------------------------------------------------

type Inserter[T Identifiable] interface {
	Insert(t T) Error
}

func (d *database[T]) Insert(t T) Error {
	d.eventChan <- NewEvent(eventTypeInsert, t)
	return nil
}
