package validator

import "github.com/go-playground/validator/v10"

type ValidatorUtil interface {
	Validate(param any) error
}

type validatorUtil struct {
	validator *validator.Validate
}

func NewValidatorUtil() ValidatorUtil {
	return &validatorUtil{
		validator: validator.New(),
	}
}

func (v *validatorUtil) Validate(param any) error {
	return v.validator.Struct(param)
}
