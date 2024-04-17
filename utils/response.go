package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response struct for standardizing JSON responses
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error struct for standardizing JSON error responses
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessResponse function to create a successful response
func SuccessResponse(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Success"
	}

	return JSONResponse(c, http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse function to create an error response
func ErrorResponse(c echo.Context, code int, message string) error {
	return JSONResponse(c, code, Error{
		Code:    code,
		Message: message,
	})
}

// JSONResponse function to create a JSON response
func JSONResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}
