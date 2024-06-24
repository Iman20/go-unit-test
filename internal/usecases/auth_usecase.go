package usecases

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-unit-test/internal/models"
	"go-unit-test/internal/repositories"
	"go-unit-test/internal/utils"
	bcrypt2 "go-unit-test/internal/utils/commons/bcrypt"
	jwt2 "go-unit-test/internal/utils/commons/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUseCase interface {
	SignUp(username, password, email string) error
	SignIn(username, password string) (string, error)
}

type authUseCase struct {
	authRepo       repositories.AuthRepository
	passwordHasher bcrypt2.BcryptBuilder
	jwtBuilder     jwt2.JwtBuilder
}

func (a authUseCase) SignUp(username, password, email string) error {
	if len(password) == 0 {
		return errors.New("invalid password")
	}

	if !utils.ValidateEmail(email) {
		return errors.New("invalid email")
	}

	hashPassword, err := a.passwordHasher.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{Username: username, Password: string(hashPassword), Email: email}
	return a.authRepo.CreateUser(user)
}

func (a authUseCase) SignIn(username, password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("invalid password")
	}

	if len(username) == 0 {
		return "", errors.New("invalid username")
	}

	user, err := a.authRepo.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	if err := a.passwordHasher.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := a.jwtBuilder.GenerateJwt(jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewAuthUseCase(repo repositories.AuthRepository, passwordHasher bcrypt2.BcryptBuilder, jwtBuilder jwt2.JwtBuilder) AuthUseCase {
	return &authUseCase{authRepo: repo, passwordHasher: passwordHasher, jwtBuilder: jwtBuilder}
}
