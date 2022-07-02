package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"

	"github.com/micheltank/eth-fee-calculator/internal/domain"
	"github.com/micheltank/eth-fee-calculator/internal/port/rest/presenter"
)

type TransactionService interface {
	GetTransactionsPerHour(from int64, to int64) ([]domain.TransactionCostPerHour, error)
}

func MakeTransactionHandler(routerGroup gin.IRoutes, service TransactionService) {
	routerGroup.GET("/transactions/cost-per-hour", func(c *gin.Context) {
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
// @Produce json
// @Success 200 {object} presenter.TransactionCostPerHourResponse
// @Error 400 {object} presenter.ApiError
// @Router /transactions/cost-per-hour [get]
func V1GetTransactionCostsPerHour(c *gin.Context, service TransactionService) {
	from, to, err := getQueryParams(c)
	if err != nil {
		logrus.WithError(err).Warn("V1GetTransactionCostsPerHour received invalid input from client")
		c.JSON(http.StatusBadRequest, presenter.ApiError{Key: "error.validationError.request", Message: err.Error()})

		return
	}
	transactions, err := service.GetTransactionsPerHour(from, to)
	if err != nil {
		logrus.WithError(err).Error("V1GetTransactionCostsPerHour returned InternalServerError")
		c.Status(http.StatusInternalServerError)

		return
	}
	c.JSON(http.StatusOK, presenter.NewTransactionCostsPerHourResponse(transactions))
}

func getQueryParams(c *gin.Context) (int64, int64, error) {
	var result error

	from, err := getQueryParam(c, "from")
	if err != nil {
		result = multierror.Append(result, err)
	}
	to, err := getQueryParam(c, "to")
	if err != nil {
		result = multierror.Append(result, err)
	}
	return from, to, err
}

func getQueryParam(c *gin.Context, queryParamName string) (int64, error) {
	fromParam := c.Query(queryParamName)
	if fromParam == "" {
		return 0, errors.New(fmt.Sprintf("'%s' query param is missing", queryParamName))
	}
	from, err := strconv.ParseInt(fromParam, 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("query param '%s' with value '%d' has invalid type, got: '%T', expected: 'int64'", queryParamName, from, from))
	}
	return from, nil
}
