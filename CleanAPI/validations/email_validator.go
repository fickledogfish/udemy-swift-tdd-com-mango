package validations

import "net/mail"

type emailValidator struct {
}

func NewEmailValidator() emailValidator {
	return emailValidator{}
}

// Implementing Validation ----------------------------------------------------

func (v emailValidator) IsValid(email any) (errs []Error) {
	emailString, ok := email.(string)
	if !ok {
		return []Error{VErrInvalidData{}}
	}

	_, err := mail.ParseAddress(emailString)
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
