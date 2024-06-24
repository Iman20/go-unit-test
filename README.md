# Go Unit Test

This directory contains unit tests for the `Go Unit Test` using Ginkgo and Gomega.

## Running the Tests

### Prerequisites

Ensure you have Ginkgo and Gomega installed. You can install them using:

```bash
go install github.com/onsi/ginkgo/v2/ginkgo@latest
go get -u github.com/onsi/gomega/...
```

To run the tests, navigate to the directory containing the tests and run:

```bash
ginkgo -vv
```

To run all tests in the project recursively from the root directory:
```bash
ginkgo -r -vv
```

Run Tests with JUnit Report:
```bash
ginkgo -r -vv -junit-report report.xml -output-dir ./test-reports
```

## Documentation API

Using swagger go, install package:

```bash
go get -u github.com/labstack/echo/v4
go get -u github.com/swaggo/echo-swagger
```

Generate Swagger Documentation:

```bash
swag init
```

Import echo and echo-swagger correctly and set up the swagger route

```go
package main

import (
    ...
    echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Auth API
// @version 1.0
// @description This is a sample server for an authentication API.
// @host localhost:8080
// @BasePath /
func main() {
    e := echo.New()
	...
    // Swagger route
    e.GET("/swagger/*", echoSwagger.WrapHandler)
	...
}
```
Create annotation in controller
```go
// @Summary Sign up a new user
// @Description Register a new user with username, password, and email
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username body string true "Username"
// @Param   password body string true "Password"
// @Param   email body string true "Email"
// @Success 200 {object} response.SuccessResponse "User signed up successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid input"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /signup [post]
func (a *AuthController) SignUp(c echo.Context) error {
	...
	return nil
}
```