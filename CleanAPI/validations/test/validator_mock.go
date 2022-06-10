package validationtest

import (
	v "example.com/api/validations"
)

type validateFunc[T any] func(T) []v.Error

type ValidatorMock[T any] struct {
	ValidateWith validateFunc[T]
}

// Implementing Validation ----------------------------------------------------

func (m ValidatorMock[T]) Validate(data T) []v.Error {
	if m.ValidateWith != nil {
		return m.ValidateWith(data)
	} else {
		return []v.Error{}
	}
}
