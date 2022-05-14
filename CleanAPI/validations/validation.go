package validations

type Validation[T any] interface {
	IsValid(T) []Error
}

type Error interface {
	error
}
