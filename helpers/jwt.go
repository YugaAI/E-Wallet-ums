package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 1,
	"refresh_token": time.Hour * 200,
}

var jwtSecret = []byte(os.Getenv("APP_SECRET"))

func GenerateToken(ctx context.Context, userID int, username, fullname, tokenType string, now time.Time) (string, error) {

	claimToken := ClaimToken{
		UserID:   userID,
		Username: username,
		Fullname: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	resutToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return resutToken, fmt.Errorf("failed to generate token: %w", err)
	}
	return resutToken, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {

	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	claimToken, ok := jwtToken.Claims.(*ClaimToken)
	if !ok || !jwtToken.Valid {
		return claimToken, fmt.Errorf("invalid token")
	}
	return claimToken, nil
}
