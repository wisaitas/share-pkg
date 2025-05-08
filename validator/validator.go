package validator

import "github.com/go-playground/validator/v10"

type Util interface {
	Validate(v any) error
}

type util struct {
	validate *validator.Validate
}

func New() Util {
	return &util{validate: validator.New()}
}

func (r *util) Validate(v any) error {
	return r.validate.Struct(v)
}
