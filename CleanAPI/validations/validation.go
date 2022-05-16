package validations

type Validation[T any] interface {
	Validate(T) []Error
}

type Error interface {
	error
}
