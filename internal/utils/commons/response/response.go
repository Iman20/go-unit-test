package response

import "github.com/labstack/echo/v4"

type MapResponse map[string]interface{}

func Error(r echo.Context, statusCode int, err error) {
	r.JSON(statusCode, MapResponse{
		"error_code":        statusCode,
		"error_description": err.Error(),
	})
}

func Success(r echo.Context, statusCode int, data interface{}) {
	r.JSON(statusCode, MapResponse{
		"status_code": statusCode,
		"message":     "success",
		"data":        data,
	})
}
