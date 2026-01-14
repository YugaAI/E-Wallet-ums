package interfaces

import "github.com/gin-gonic/gin"

type IHealtCheckServices interface {
	HealtCheckServices() (string, error)
}
type IHealtCheckHandler interface {
	HealthCheckHandlerHTTP(c *gin.Context)
}
type IHealthCheckRepository interface {
}
