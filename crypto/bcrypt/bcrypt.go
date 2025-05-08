package bcrypt

import "golang.org/x/crypto/bcrypt"

type Util interface {
	GenerateFromPassword(password string, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type util struct {
}

func New() Util {
	return &util{}
}

func (r *util) GenerateFromPassword(password string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

func (r *util) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
