package repository

import (
	"context"
	"errors"
	"ewallet-framework/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertUser(ctx context.Context, user *model.Users) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (model.Users, error) {
	var (
		err  error
		user model.Users
	)
	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *model.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) DeleteNewUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_session WHERE token = ?", token).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (model.UserSession, error) {
	var (
		err       error
		userToken model.UserSession
	)
	err = r.DB.Where("token = ?", token).First(&userToken).Error
	if err != nil {
		return userToken, err
	}
	if userToken.ID == 0 {
		return userToken, errors.New("user not found")
	}
	return userToken, nil
}
