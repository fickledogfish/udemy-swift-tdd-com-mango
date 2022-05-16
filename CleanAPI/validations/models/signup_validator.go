package models

import (
	m "example.com/api/models"
	v "example.com/api/validations"
)

type SignUpModelValidator struct {
	emailValidator v.Validation[string]
}

func NewSignUpModelValidator(
	emailValidator v.Validation[string],
) SignUpModelValidator {
	return SignUpModelValidator{
		emailValidator: emailValidator,
	}
}

// Implementing Validation ----------------------------------------------------

func (v SignUpModelValidator) Validate(data m.SignUp) (errors []v.Error) {
	errors = append(errors, v.emailValidator.Validate(data.Email)...)

	return
}
