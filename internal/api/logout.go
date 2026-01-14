package api

import (
	"ewallet-framework/constants"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutSvc interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.GetHeader("Authorization")
	if token == "" {
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		return
	}
	
	err := api.LogoutSvc.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("Faild to Logout", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, nil)

}
