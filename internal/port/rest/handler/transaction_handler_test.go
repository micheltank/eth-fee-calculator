package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"

	"github.com/micheltank/eth-fee-calculator/internal/domain"
	transactionHandler "github.com/micheltank/eth-fee-calculator/internal/port/rest/handler"
	mockApplication "github.com/micheltank/eth-fee-calculator/internal/port/rest/handler/mock"
	"github.com/micheltank/eth-fee-calculator/internal/port/rest/presenter"
)

func TestV1GetTransactionCostsPerHour(t *testing.T) {

	t.Run("Regular get", func(t *testing.T) {
		g := NewGomegaWithT(t)

		var from int64 = 1599476913
		var to int64 = 1599649713

		transactions := []domain.TransactionCostPerHour{
			{
				Hour:      time.Unix(from, 0),
				FeeAmount: 123.45,
			},
			{
				Hour:      time.Unix(from+int64(time.Hour), 0),
				FeeAmount: 456.78,
			},
		}
		service := mockService(t, from, to, transactions, nil)

		route := "/transactions/cost-per-hour"
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET(route, func(c *gin.Context) {
			transactionHandler.V1GetTransactionCostsPerHour(c, service)
		})
		routeWithParams := fmt.Sprintf(route+"?from=%d&to=%d", from, to)
		req, err := http.NewRequest("GET", routeWithParams, nil)
		assert.NoError(t, err)

		r.ServeHTTP(w, req)

		g.Expect(w.Code).Should(
			Equal(http.StatusOK))

		var got []presenter.TransactionCostPerHourResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		assert.NoError(t, err)

		responseExpect := presenter.NewTransactionCostsPerHourResponse(transactions)
		g.Expect(responseExpect).Should(
			Equal(got))
	})

	t.Run("Get with invalid query params", func(t *testing.T) {
		g := NewGomegaWithT(t)

		var from int64 = 1599476913
		var to int64 = 1599649713

		var transactions []domain.TransactionCostPerHour
		service := mockService(t, from, to, transactions, nil)

		route := "/transactions/cost-per-hour"
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		transactionHandler.MakeTransactionHandler(r, service)

		routeWithParams := fmt.Sprintf(route + "?from=asdad&to=asdasd")
		req, err := http.NewRequest("GET", routeWithParams, nil)
		assert.NoError(t, err)

		r.ServeHTTP(w, req)

		g.Expect(w.Code).Should(
			Equal(http.StatusBadRequest))

		var got presenter.ApiError
		err = json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		apiError := presenter.ApiError{
			Key: "error.validationError.request",
		}
		g.Expect(apiError.Key).Should(
			Equal(got.Key))
	})
	t.Run("Get transaction with absent query param", func(t *testing.T) {
		g := NewGomegaWithT(t)

		var from int64 = 1599476913
		var to int64 = 1599649713

		var transactions []domain.TransactionCostPerHour
		service := mockService(t, from, to, transactions, nil)

		route := "/transactions/cost-per-hour"
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET(route, func(c *gin.Context) {
			transactionHandler.V1GetTransactionCostsPerHour(c, service)
		})
		req, err := http.NewRequest("GET", route, nil)
		assert.NoError(t, err)

		r.ServeHTTP(w, req)

		g.Expect(w.Code).Should(
			Equal(http.StatusBadRequest))

		var got presenter.ApiError
		err = json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		apiError := presenter.ApiError{
			Key: "error.validationError.request",
		}
		g.Expect(apiError.Key).Should(
			Equal(got.Key))
	})
	t.Run("Get transaction with service error", func(t *testing.T) {
		g := NewGomegaWithT(t)

		var from int64 = 1599476913
		var to int64 = 1599649713

		var transactions []domain.TransactionCostPerHour
		service := mockService(t, from, to, transactions, errors.New("dummy error"))

		route := "/transactions/cost-per-hour"
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET(route, func(c *gin.Context) {
			transactionHandler.V1GetTransactionCostsPerHour(c, service)
		})
		routeWithParams := fmt.Sprintf(route+"?from=%d&to=%d", from, to)
		req, err := http.NewRequest("GET", routeWithParams, nil)
		assert.NoError(t, err)

		r.ServeHTTP(w, req)

		g.Expect(w.Code).Should(
			Equal(http.StatusInternalServerError))

		g.Expect(w.Body.Len()).Should(
			Equal(0))
	})
}

func mockService(t *testing.T, from, to int64, transactionResponse []domain.TransactionCostPerHour, errorResponse error) *mockApplication.MockTransactionService {
	ctrl := gomock.NewController(t)
	service := mockApplication.NewMockTransactionService(ctrl)
	service.EXPECT().
		GetTransactionsPerHour(gomock.Eq(from), gomock.Eq(to)).
		DoAndReturn(func(from, to int64) ([]domain.TransactionCostPerHour, error) {
			return transactionResponse, errorResponse
		}).
		AnyTimes()
	return service
}
