package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-unit-test/internal/usecases"
	"go-unit-test/internal/utils/commons/response"
	"net/http"
)

type ControllerInterface interface {
	SignUp(ctx echo.Context) error
	SignIn(ctx echo.Context) error
}

type controllerImpl struct {
	authCase usecases.AuthUseCase
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @Summary Sign up a new user
// @Description Register a new user with username, password, and email
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username body string true "Username"
// @Param   password body string true "Password"
// @Param   email body string true "Email"
// @Success 200 {string} string "User signed up successfully"
// @Router /v1/user/signup [post]
func (c controllerImpl) SignUp(ctx echo.Context) error {
	req := new(SignUpRequest)
	if err := ctx.Bind(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, errors.New("invalid request"))
		return nil
	}

	err := c.authCase.SignUp(req.Username, req.Password, req.Email)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, err)
		return nil
	}

	response.Success(ctx, http.StatusOK, nil)
	return nil
}

// @Summary Sign in a user
// @Description Sign in a user with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username body string true "Username"
// @Param   password body string true "Password"
// @Success 200 {string} string "User signed up successfully"
// @Router /v1/user/signin [post]
func (c controllerImpl) SignIn(ctx echo.Context) error {
	req := new(SignInRequest)
	if err := ctx.Bind(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, errors.New("invalid request"))
		return nil
	}
	token, err := c.authCase.SignIn(req.Username, req.Password)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, err)
		return nil
	}

	response.Success(ctx, http.StatusOK, map[string]string{"token": token})
	return nil
}

func NewController(authCase usecases.AuthUseCase) ControllerInterface {
	return &controllerImpl{authCase: authCase}
}
