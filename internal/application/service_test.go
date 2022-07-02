package application_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/micheltank/eth-fee-calculator/internal/application"
	"github.com/micheltank/eth-fee-calculator/internal/application/mock"
	"github.com/micheltank/eth-fee-calculator/internal/domain"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mock.NewMockRepository(ctrl)

	var from int64 = 1599476913
	var to int64 = 1599649713
	transactionsPerHourExpected := []domain.TransactionCostPerHour{
		{
			Hour:      time.Unix(from, 0),
			FeeAmount: 123.45,
		},
		{
			Hour:      time.Unix(from+int64(time.Hour), 0),
			FeeAmount: 456.78,
		},
	}
	errorExpected := errors.New("dummy error")

	repository.EXPECT().
		GetTransactionsPerHour(gomock.Eq(from), gomock.Eq(to), 1).
		DoAndReturn(func(from, to int64, page int) ([]domain.TransactionCostPerHour, error) {
			return transactionsPerHourExpected, nil
		}).
		AnyTimes()

	repository.EXPECT().
		GetTransactionsPerHour(gomock.Eq(int64(1)), gomock.Eq(int64(2)), 1).
		DoAndReturn(func(from, to int64, page int) ([]domain.TransactionCostPerHour, error) {
			return nil, errorExpected
		}).
		AnyTimes()

	t.Run("Regular get transactions per hour", func(t *testing.T) {
		g := NewGomegaWithT(t)

		service := application.NewService(repository)

		transactionsPerHour, err := service.GetTransactionsPerHour(from, to, 1)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(transactionsPerHour).Should(
			Not(BeNil()))
		g.Expect(len(transactionsPerHour)).Should(
			Equal(len(transactionsPerHourExpected)))
	})

	t.Run("Get transactions per hour with repository error", func(t *testing.T) {
		g := NewGomegaWithT(t)

		service := application.NewService(repository)

		transactionsPerHour, err := service.GetTransactionsPerHour(int64(1), int64(2), 1)
		g.Expect(errorExpected).Should(MatchError(errors.Cause(err)))
		g.Expect(transactionsPerHour).Should(
			BeEmpty())
	})
}
