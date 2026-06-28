package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse[T any] struct {
	Success bool      `json:"success"`
	Data    T         `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
}

func SuccessResponse[T any](c *gin.Context, statusCode int, data T) {
	c.JSON(statusCode, APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func OKResponse[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, APIResponse[any]{
		Success: false,
		Error: &APIError{
			Code:    statusCode,
			Message: message,
		},
	})
}
