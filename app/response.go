package app

import (
	"github.com/gin-gonic/gin"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status  Status      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func ErrorResponse(c *gin.Context, code int, message string, errorMessage string) {
	if errorMessage == "" {
		errorMessage = "Error"
	}
	response := Response{
		Status: Status{
			Code:    code,
			Message: errorMessage,
		},
		Message: message,
	}
	c.JSON(code, response)
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	response := Response{
		Status: Status{
			Code:    code,
			Message: "Success",
		},
		Message: message,
		Data:    data,
	}
	c.JSON(code, response)
}
