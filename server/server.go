package server

import (
	"github.com/labstack/echo/v4"
	"go-unit-test/config"
	"go-unit-test/internal/controllers"
	"go-unit-test/internal/repositories"
	"go-unit-test/internal/usecases"
	"go-unit-test/internal/utils/commons/bcrypt"
	"go-unit-test/internal/utils/commons/jwt"
)

func RegisterService(r *echo.Echo) {
	db := config.InitDB()

	passwordHasher := bcrypt.NewBycryptBuilder()
	jwtBuilder := jwt.NewJwtBuilder()

	authRepo := repositories.NewAuthRepository(db)
	useCase := usecases.NewAuthUseCase(authRepo, passwordHasher, jwtBuilder)
	controller := controllers.NewController(useCase)

	registerRouter(r, controller)
}
