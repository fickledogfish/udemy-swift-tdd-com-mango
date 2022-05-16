package validations

import "net/mail"

type EmailValidator struct {
}

func NewEmailValidator() EmailValidator {
	return EmailValidator{}
}

// Implementing Validation ----------------------------------------------------

func (v EmailValidator) Validate(email string) (errs []Error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		errs = append(errs, VErrInvalidEmail{})
	}

	return
}

// Email validation errors ----------------------------------------------------

type VErrInvalidEmail struct{}

func (e VErrInvalidEmail) Error() string {
	return "Invalid email address"
}
