package models

import (
	m "example.com/api/models"
	vs "example.com/api/validations"
)

type SignUpModelValidator struct {
	emailValidator    vs.Validation[string]
	passwordValidator vs.Validation[vs.PasswordValidatorData]
}

func NewSignUpModelValidator(
	emailValidator vs.Validation[string],
	passwordValidator vs.Validation[vs.PasswordValidatorData],
) SignUpModelValidator {
	return SignUpModelValidator{
		emailValidator:    emailValidator,
		passwordValidator: passwordValidator,
	}
}

// Implementing Validation ----------------------------------------------------

func (v SignUpModelValidator) Validate(data m.SignUp) (errors []vs.Error) {
	errors = append(errors, v.emailValidator.Validate(data.Email)...)
	errors = append(errors, v.passwordValidator.Validate(
		vs.PasswordValidatorData{
			Password:     data.Password,
			Confirmation: data.PasswordConfirmation,
		},
	)...)

	return
}
