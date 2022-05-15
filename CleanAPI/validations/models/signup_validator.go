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

func (v SignUpModelValidator) IsValid(data m.SignUp) (errors []v.Error) {
	errors = append(errors, v.emailValidator.IsValid(data.Email)...)

	return
}
