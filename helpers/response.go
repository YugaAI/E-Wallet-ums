package helpers

import (
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	resp := HealthResponse{
		Message: message,
		Data:    data,
	}
	c.JSON(code, resp)
}
