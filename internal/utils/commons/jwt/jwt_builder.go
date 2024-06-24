package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

type JwtBuilder interface {
	GenerateJwt(claim jwt.Claims) (string, error)
}

type JwtBuilderImpl struct{}

func (j JwtBuilderImpl) GenerateJwt(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(jwtSecret)
}

func NewJwtBuilder() JwtBuilder {
	return JwtBuilderImpl{}
}
