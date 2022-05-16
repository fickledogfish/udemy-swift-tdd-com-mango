package validationtest

import v "example.com/api/validations"

type validateFunc func(string) []v.Error

type emailValidatorMock struct {
	validateWith validateFunc
}

// Creates a new email validator mock with the specified callback, which is
// executed when Validate(string) is called.
func NewEmailValidatorMock(validateWith validateFunc) emailValidatorMock {
	return emailValidatorMock{
		validateWith: validateWith,
	}
}

// Implementing Validation ----------------------------------------------------

func (v emailValidatorMock) Validate(email string) []v.Error {
	return v.validateWith(email)
}
