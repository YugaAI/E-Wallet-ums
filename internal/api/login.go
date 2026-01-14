package api

import (
	"ewallet-framework/constants"
	"ewallet-framework/helpers"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginSvc interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  model.LoginRequest
		resp model.LoginResponse
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Faild to Parse Request ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}
	if err := req.Validate(); err != nil {
		log.Error("Faild to Validate Request", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}

	resp, err := api.LoginSvc.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Faild to Login", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)

}
