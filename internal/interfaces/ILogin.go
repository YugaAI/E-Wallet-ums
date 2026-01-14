package interfaces

import (
	"context"
	"ewallet-framework/internal/model"

	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}
type ILoginHandler interface {
	Login(c *gin.Context)
}
