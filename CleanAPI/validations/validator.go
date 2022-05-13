package validations

type Validation interface {
	IsValid(any) []error
}

type Error interface {
	error
}

// Validation errors ----------------------------------------------------------

type VErrInvalidData struct{}

func (e VErrInvalidData) Error() string {
	return "Invalid data"
}
