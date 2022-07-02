package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"

	"github.com/micheltank/eth-fee-calculator/internal/port/rest/handler"
)

func TestHealthCheck(t *testing.T) {

	t.Run("Health check success", func(t *testing.T) {
		g := NewGomegaWithT(t)

		route := "/health-check"
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		handler.MakeHealthCheckHandler(r)

		req, err := http.NewRequest("GET", route, nil)
		assert.NoError(t, err)

		r.ServeHTTP(w, req)

		g.Expect(w.Code).Should(
			Equal(http.StatusOK))
	})
}
