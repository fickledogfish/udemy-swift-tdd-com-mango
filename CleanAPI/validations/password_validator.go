package validations

type PasswordValidatorData struct {
	Password     string
	Confirmation string
}

type PasswordValidator struct{}

func NewPasswordValidator() PasswordValidator {
	return PasswordValidator{}
}

// Implementing Validation ----------------------------------------------------

func (v PasswordValidator) Validate(
	data PasswordValidatorData,
) (violations []Error) {
	if data.Password != data.Confirmation {
		violations = append(violations, passwordAndConfirmationMismatch())
	}

	return
}

// Password validation errors -------------------------------------------------

type VErrPasswordAndConfirmationMismatch struct{}

func passwordAndConfirmationMismatch() VErrPasswordAndConfirmationMismatch {
	return VErrPasswordAndConfirmationMismatch{}
}

func (e VErrPasswordAndConfirmationMismatch) Error() string {
	return "Password and confirmation do not match"
}
