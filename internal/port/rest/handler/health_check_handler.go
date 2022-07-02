package handler

import (
	"github.com/micheltank/eth-fee-calculator/internal/infra/config"
	"github.com/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v4"
	healthPostgres "github.com/hellofresh/health-go/v4/checks/postgres"
)

func MakeHealthCheckHandler(routerGroup gin.IRoutes, dbConfig config.DbConfig) error {
	h, err := health.New(health.WithChecks(health.Config{
		Name:      "postgres",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthPostgres.New(healthPostgres.Config{
			DSN: dbConfig.BuildURL(),
		}),
	}))
	if err != nil {
		return errors.Wrap(err, "MakeHealthCheckHandler: failed to initialize health check")
	}

	routerGroup.GET("/health-check", func(c *gin.Context) {
		h.HandlerFunc(c.Writer, c.Request)
	})

	return nil
}
