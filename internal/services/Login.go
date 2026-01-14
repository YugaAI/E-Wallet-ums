package services

import (
	"context"
	"errors"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	LoginRepo interfaces.IUserRepository
}

func (svc *LoginService) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	var (
		resp model.LoginResponse
		now  = time.Now()
	)
	userDetail, err := svc.LoginRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return resp, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return resp, errors.New("password incorrect")
	}

	token, errToken := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if errToken != nil {
		return resp, errors.New("token generation failed")
	}

	refreshToken, errRefreshToken := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", now)
	if errRefreshToken != nil {
		return resp, errors.New("Refresh Token generation failed")
	}
	userSession := &model.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		TokenRefreshExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = svc.LoginRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.New(" user session creation failed")
	}

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Token = token
	resp.RefreshToken = refreshToken

	return resp, nil
}
