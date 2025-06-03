package validatorutil

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

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
	if err := v.validator.Struct(param); err != nil {
		return fmt.Errorf("[Share Package ValidatorUtil] : %w", err)
	}

	return nil
}
