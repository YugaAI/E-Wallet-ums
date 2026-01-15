package interfaces

import (
	"context"
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
)

type ITokenValidateService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

type ITokenValidateHandler interface {
	ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}
