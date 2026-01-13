package interfaces

import (
	"context"
	"ewallet-framework/internal/model"
)

type IRegisterRepository interface {
	InsertUser(ctx context.Context, user *model.Users) error
}

type IRegisterService interface {
	Register(ctx context.Context, request model.Users) (interface{}, error)
}
