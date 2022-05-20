package validationtest

import (
	v "example.com/api/validations"
)

type validateFunc[T any] func(T) []v.Error

type ValidatorMock[T any] struct {
	ValidateWith validateFunc[T]
}

func NewValidatorMock[T any](validateWith validateFunc[T]) ValidatorMock[T] {
	return ValidatorMock[T]{
		ValidateWith: validateWith,
	}
}

// Implementing Validation ----------------------------------------------------

func (v ValidatorMock[T]) Validate(data T) []v.Error {
	return v.ValidateWith(data)
}
