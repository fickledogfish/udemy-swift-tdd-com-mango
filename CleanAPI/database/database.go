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
	ch := make(chan event[T])

	go func(ch <-chan event[T]) {
		for event := range ch {
			fmt.Println(event)
		}
	}(ch)

	return &database[T]{
		Data: []T{},

		eventChan: ch,
	}
}

// Implementing Inserter ------------------------------------------------------

type Inserter[T Identifiable] interface {
	Insert(t T)
}

func (d *database[T]) Insert(t T) {
}
