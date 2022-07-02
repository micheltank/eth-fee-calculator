package rest

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/micheltank/eth-fee-calculator/internal/infra/config"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/micheltank/eth-fee-calculator/internal/application"
	"github.com/micheltank/eth-fee-calculator/internal/infra/repository"
	"github.com/micheltank/eth-fee-calculator/internal/port/rest/handler"
)

type Api struct {
	httpServer *http.Server
	Db         *sql.DB
}

// NewServer godoc
// @title ETH Fee Calculator
// @BasePath /api/v1
// @version 1.0
func NewServer(config config.Config) (*Api, error) {
	router := gin.Default()
	base := router.Group("/api")

	v1 := base.Group("/v1")

	// di
	db, err := sql.Open("postgres", config.DbConfig.BuildURL())
	if err != nil {
		return nil, errors.Wrap(err, "NewServer: failed to open postgres connection")
	}
	transactionsRepository := repository.NewTransactionPostgreSql(db)
	service := application.NewService(transactionsRepository)

	// handlers
	err = handler.MakeHealthCheckHandler(base, config.DbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "NewServer: failed to create health check handler")
	}
	handler.MakeTransactionHandler(v1, service)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: router}

	return &Api{
		httpServer: httpServer,
	}, nil
}

func (api *Api) Run() <-chan error {
	out := make(chan error)
	go func() {
		if err := api.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			out <- errors.Wrap(err, "Run: failed to listen and serve api")
		}
	}()
	return out
}

func (api *Api) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := api.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.
				WithError(err).
				Error("Server forced to shutdown")
		}
	}()
}
