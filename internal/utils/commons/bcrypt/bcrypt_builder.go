package bcrypt

import "golang.org/x/crypto/bcrypt"

type BcryptBuilder interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}

type BycryptBuilderImpl struct{}

func (h BycryptBuilderImpl) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (BycryptBuilderImpl) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func NewBycryptBuilder() BcryptBuilder {
	return BycryptBuilderImpl{}
}
