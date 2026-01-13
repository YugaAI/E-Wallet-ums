package helpers

import (
	"ewallet-framework/internal/model"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySQL() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", ""),
		GetEnv("DB_PORT", ""),
		GetEnv("DB_NAME", ""),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Field connecting to database", err)
	}
	logrus.Info("Connected to database")

	DB.AutoMigrate(
		&model.Users{},
		&model.UserSession{},
	)
}
