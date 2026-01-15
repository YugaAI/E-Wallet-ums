package interfaces

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/model"

	"github.com/gin-gonic/gin"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, tokenClaim helpers.ClaimToken, refreshToken string) (model.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}
