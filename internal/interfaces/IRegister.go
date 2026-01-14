package interfaces

import (
	"context"
	"ewallet-framework/internal/model"

	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(ctx context.Context, request model.Users) (interface{}, error)
}
type IRegisterHandler interface {
	RegisterHandlerHTTP(c *gin.Context)
}
