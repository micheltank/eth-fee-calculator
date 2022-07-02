package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeHealthCheckHandler(routerGroup gin.IRoutes) {
	routerGroup.GET("/health-check", func(c *gin.Context) {
		HealthCheck(c)
	})
}

func HealthCheck(c *gin.Context) {
	// TODO: add db
	c.Writer.WriteHeader(http.StatusOK)
}
