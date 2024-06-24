package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-unit-test/internal/controllers"

	_ "go-unit-test/docs" // This is important for the swag init
)

func registerRouter(r *echo.Echo, controller controllers.ControllerInterface) {
	otpRouter := r.Group("/v1/user")
	otpRouter.POST("/signin", controller.SignIn)
	otpRouter.POST("/signup", controller.SignUp)
	otpRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
