package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
)

type LogoutService struct {
	LogoutRepo interfaces.IUserRepository
}

func (svc *LogoutService) Logout(ctx context.Context, token string) error {
	return svc.LogoutRepo.DeleteNewUserSession(ctx, token)
}
