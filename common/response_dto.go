package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseDTO struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func GenerateSuccessResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, ResponseDTO{
		Data:    data,
		Message: message,
		Status:  "success",
	})
}

func GenerateErrorResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusBadRequest, ResponseDTO{
		Data:    data,
		Message: message,
		Status:  "error",
	})
}

func GenerateNotFoundResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusNoContent, ResponseDTO{
		Data:    data,
		Message: message,
		Status:  "error",
	})
}
