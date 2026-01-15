package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenSvc interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := c.Request.Header.Get("Authorization")
	if refreshToken == "" {
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		return
	}
	claim, ok := c.Get("token")
	if !ok {
		log.Error("Faild to get claim in context")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("Faild to parse token claim in context")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	resp, err := api.RefreshTokenSvc.RefreshToken(c.Request.Context(), *tokenClaim, refreshToken)
	if err != nil {
		log.Error("Faild to Refresh Token", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}
