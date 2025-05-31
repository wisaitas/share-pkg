package validatorutil

import "github.com/go-playground/validator/v10"

type ValidatorUtil interface {
	ValidateStruct(param any) error
}

type validatorUtil struct {
	validator *validator.Validate
}

func New() ValidatorUtil {
	return &validatorUtil{validator: validator.New()}
}

func (v *validatorUtil) ValidateStruct(param any) error {
	return v.validator.Struct(param)
}
