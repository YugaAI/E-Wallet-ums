package cmd

import (
	"ewallet-framework/constants"
	"ewallet-framework/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(c *gin.Context) {
	if d == nil || d.UserRepository == nil {
		log.Println("UserRepository is nil")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		c.Abort()
		return
	}

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("auth is empty")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(c.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}
	claim, err := helpers.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("token is expired")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}
	c.Set("token", claim)

	c.Next()
	return
}

func (d *Dependency) MiddlewareLogout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}

	// CUMA cek DB
	_, err := d.UserRepository.GetUserSessionByToken(c.Request.Context(), token)
	if err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}

	c.Next()
}
