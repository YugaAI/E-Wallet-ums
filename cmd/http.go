package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/interfaces"
	"ewallet-framework/internal/repository"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	dependency := dependencyInject()
	r := gin.Default()

	r.GET("/health", dependency.HealthcheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.RegisterHandlerHTTP)
	userV1.POST("/login", dependency.LoginApi.Login)
	userV1.DELETE("/logout", dependency.MiddlewareLogout, dependency.LogoutApi.Logout)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository
	HealthcheckAPI interfaces.IHealtCheckHandler
	RegisterAPI    interfaces.IRegisterHandler
	LoginApi       interfaces.ILoginHandler
	LogoutApi      interfaces.ILogoutHandler
}

func dependencyInject() Dependency {
	healthCheckSvc := &services.HealthCheck{}
	healthcheckAPI := api.HealthCheckService{
		HealthCheckServices: healthCheckSvc,
	}
	userRepo := &repository.UserRepository{
		DB: helpers.GetDB(),
	}
	userService := &services.RegisterService{
		UserRepo: userRepo,
	}
	registerAPI := api.RegisterAPI{
		RegisterSvc: userService,
	}
	loginService := &services.LoginService{
		LoginRepo: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginSvc: loginService,
	}
	logoutSvc := &services.LogoutService{
		LogoutRepo: userRepo,
	}
	logoutAPI := &api.LogoutHandler{
		LogoutSvc: logoutSvc,
	}
	return Dependency{
		UserRepository: userRepo,
		HealthcheckAPI: &healthcheckAPI,
		RegisterAPI:    &registerAPI,
		LoginApi:       loginAPI,
		LogoutApi:      logoutAPI,
	}
}
