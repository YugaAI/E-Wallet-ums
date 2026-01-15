package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
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

	userV1WithAuth := userV1.Use()
	userV1WithAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutApi.Logout)
	userV1WithAuth.PUT("/refresh-token", dependency.MiddlewareRefreshToken, dependency.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository     interfaces.IUserRepository
	HealthcheckAPI     interfaces.IHealtCheckHandler
	RegisterAPI        interfaces.IRegisterHandler
	LoginApi           interfaces.ILoginHandler
	LogoutApi          interfaces.ILogoutHandler
	RefreshTokenAPI    interfaces.IRefreshTokenHandler
	ValidationTokenAPI *api.TokenValidationAPI
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
	RefreshTokenSvc := &services.RefreshTokenService{
		RefreshTokenRepo: userRepo,
	}
	RefreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenSvc: RefreshTokenSvc,
	}
	TokenValidationSvc := &services.TokenValidationService{
		ValidateTokenRepo: userRepo,
	}
	TokenValidationAPI := &api.TokenValidationAPI{
		TokenValidationSVC: TokenValidationSvc,
	}
	return Dependency{
		UserRepository:     userRepo,
		HealthcheckAPI:     &healthcheckAPI,
		RegisterAPI:        &registerAPI,
		LoginApi:           loginAPI,
		LogoutApi:          logoutAPI,
		RefreshTokenAPI:    RefreshTokenAPI,
		ValidationTokenAPI: TokenValidationAPI,
	}
}
