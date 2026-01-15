package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/model"
	"log"
	"time"

	"github.com/pkg/errors"
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

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", userDetail.Email, now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &model.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = svc.LoginRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.New(" user session creation failed")
	}
	log.Println("LOGIN TOKEN:", token)
	log.Println("LOGIN REFRESH:", refreshToken)

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken

	return resp, nil
}
