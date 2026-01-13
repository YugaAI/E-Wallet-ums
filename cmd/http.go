package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/repository"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {

	HealthCheckSvc := &services.HealthCheck{}
	healthcheckAPI := api.HealthCheckService{
		HealthCheckServices: HealthCheckSvc,
	}
	registerRepo := &repository.RegisterRepository{
		DB: helpers.GetDB(),
	}
	registerService := &services.RegisterService{
		RegisterRepo: registerRepo,
	}
	registerAPI := api.RegisterAPI{
		RegisterSvc: registerService,
	}
	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerAPI.RegisterHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
