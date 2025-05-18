package validatorutil

import "github.com/go-playground/validator/v10"

type ValidateUtil interface {
	ValidateStruct(param any) error
}

type validateUtil struct {
	validator *validator.Validate
}

func New() ValidateUtil {
	return &validateUtil{validator: validator.New()}
}

func (r *validateUtil) ValidateStruct(param any) error {
	return r.validator.Struct(param)
}
