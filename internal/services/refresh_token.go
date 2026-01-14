package services

import (
	"context"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/model"
	"time"

	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	RefreshTokenRepo interfaces.IUserRepository
}

func (svc *RefreshTokenService) RefreshToken(ctx context.Context, tokenClaim helpers.ClaimToken, refreshToken string) (model.RefreshTokenResponse, error) {
	resp := model.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "token", tokenClaim.Email, time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate new token")
	}

	err = svc.RefreshTokenRepo.UpdateToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update new token")
	}
	resp.Token = token
	return resp, nil
}
