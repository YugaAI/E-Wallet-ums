package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
	External interfaces.IExternal
}

func (svc *RegisterService) Register(ctx context.Context, request model.Users) (interface{}, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPass)

	err = svc.UserRepo.InsertUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	_, err = svc.External.CreateWallet(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := svc.External.NotifyUserRegistered(request.ID, request.Email, request.FullName); err != nil {
		return nil, fmt.Errorf("failed to notify user: %w", err)
	}

	resp := request
	resp.Password = ""
	return resp, nil
}
