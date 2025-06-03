package bcryptutil

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptUtil interface {
	GenerateFromPassword(password string, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type bcryptUtil struct {
}

func NewBcryptUtil() BcryptUtil {
	return &bcryptUtil{}
}

func (r *bcryptUtil) GenerateFromPassword(password string, cost int) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return nil, fmt.Errorf("[Share Package BcryptUtil] : %w", err)
	}

	return hashedPassword, nil
}

func (r *bcryptUtil) CompareHashAndPassword(hashedPassword, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return fmt.Errorf("[Share Package BcryptUtil] : %w", err)
	}

	return nil
}
