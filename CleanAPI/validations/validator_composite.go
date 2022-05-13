package validations

type ValidatorComposite struct {
	validators []Validation
}

func NewValidatorComposite() ValidatorComposite {
	return ValidatorComposite{
		validators: []Validation{},
	}
}
