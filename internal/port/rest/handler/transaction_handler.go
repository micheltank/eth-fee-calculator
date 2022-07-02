package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/micheltank/eth-fee-calculator/internal/domain"
	"github.com/micheltank/eth-fee-calculator/internal/port/rest/presenter"
)

type TransactionService interface {
	GetTransactionsPerHour(from, to int64, page int) ([]domain.TransactionCostPerHour, error)
}

func MakeTransactionHandler(routerGroup gin.IRoutes, service TransactionService) {
	routerGroup.GET("/transactions/cost-per-hour", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		V1GetTransactionCostsPerHour(c, service)
	})
}

// V1GetTransactionCostsPerHour godoc
// @Summary Get transaction costs per hour
// @Description Get transaction costs per hour
// @ID get-transaction-per-hour
// @Tags Transactions
// @Param from query string true "Initial period to fetch the data"
// @Param to query string true "Final period to fetch the data"
// @Param page query string true "Pagination"
// @Produce json
// @Success 200 {object} presenter.TransactionCostPerHourResponse
// @Error 400 {object} presenter.ApiError
// @Router /transactions/cost-per-hour [get]
func V1GetTransactionCostsPerHour(c *gin.Context, service TransactionService) {
	params, err := getQueryParams(c)
	if err != nil {
		logrus.WithError(err).Warn("V1GetTransactionCostsPerHour received invalid input from client")
		c.JSON(http.StatusBadRequest, presenter.ApiError{Key: "error.validation.request", Message: err.Error()})

		return
	}
	transactions, err := service.GetTransactionsPerHour(params.From, params.To, params.Page)
	if err != nil {
		logrus.WithError(err).Error("V1GetTransactionCostsPerHour returned InternalServerError")
		c.Status(http.StatusInternalServerError)

		return
	}
	c.JSON(http.StatusOK, presenter.NewTransactionCostsPerHourResponse(transactions))
}

func getQueryParams(c *gin.Context) (presenter.TransactionCostsPerHourParams, error) {
	var params presenter.TransactionCostsPerHourParams
	err := c.BindQuery(&params)
	if err != nil {
		return params, err
	}

	err = validator.New().Struct(params)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return presenter.TransactionCostsPerHourParams{}, validationErrors
	}

	return params, nil
}
