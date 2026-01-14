package interfaces

import (
	"context"
	"ewallet-framework/internal/model"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *model.Users) error
	InsertNewUserSession(ctx context.Context, session *model.UserSession) error
	FindByUsername(ctx context.Context, username string) (model.Users, error)
	DeleteNewUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (model.UserSession, error)
}

type IRegisterService interface {
	Register(ctx context.Context, request model.Users) (interface{}, error)
}
type IRegisterHandler interface {
	RegisterHandlerHTTP(c *gin.Context)
}

type ILoginService interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}
type ILoginHandler interface {
	Login(c *gin.Context)
}

type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}
type ILogoutHandler interface {
	Logout(c *gin.Context)
}
