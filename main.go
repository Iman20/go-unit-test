package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-unit-test/server"

	_ "go-unit-test/docs" // This is important for the swag init
)

// @title Auth API
// @version 1.0
// @description This is a sample server for an authentication API.
// @host localhost:8080
// @BasePath /
func main() {
	r := echo.New()
	server.RegisterService(r)

	// Swagger route
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	r.Start(":8080")

}
