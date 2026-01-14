package interfaces

import (
	"context"
	"ewallet-framework/internal/model"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *model.Users) error
	InsertNewUserSession(ctx context.Context, session *model.UserSession) error
	FindByUsername(ctx context.Context, username string) (model.Users, error)
	DeleteNewUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (model.UserSession, error)
	UpdateToken(ctx context.Context, token, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (model.UserSession, error)
}
