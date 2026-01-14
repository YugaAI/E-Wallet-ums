package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	ID          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username    string    `json:"username" gorm:"colum:username;type:varchar(20);not null" validate:"required"`
	Email       string    `json:"email" gorm:"colum:email;type:varchar(50);not null" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"colum:phone_number;type:varchar(20);not null" validate:"required"`
	Address     string    `json:"address" gorm:"colum:address;type:text;not null"`
	DoB         string    `json:"DoB" gorm:"column:dob;type:date;not null"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);not null" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*Users) TableName() string {
	return "users"
}
func (i Users) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

type UserSession struct {
	ID                  int `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:text" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_session"
}
func (i UserSession) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}
