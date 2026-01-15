package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckService struct {
	HealthCheckServices interfaces.IHealtCheckServices
}

func (api *HealthCheckService) HealthCheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthCheckServices.HealtCheckServices()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, msg, nil)
}
