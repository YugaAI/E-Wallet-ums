package repository

import (
	"context"
	"ewallet-framework/internal/model"

	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertUser(ctx context.Context, user *model.Users) error {
	return r.DB.Create(user).Error
}
