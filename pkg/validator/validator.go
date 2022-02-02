package validator

type Validator struct {
	Errors map[string][]string
}

func New() Validator {
	return Validator{
		Errors: make(map[string][]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(field, msg string) {
	v.Errors[field] = append(v.Errors[field], msg)
}
