package application

import (
	"github.com/pkg/errors"

	"github.com/micheltank/eth-fee-calculator/internal/domain"
)

type Service struct {
	repository domain.TransactionsRepository
}

type Repository interface {
	GetTransactionsPerHour(from, to int64) ([]domain.TransactionCostPerHour, error)
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTransactionsPerHour(from, to int64) ([]domain.TransactionCostPerHour, error) {
	transactions, err := s.repository.GetTransactionsPerHour(from, to)
	if err != nil {
		return []domain.TransactionCostPerHour{}, errors.Wrap(err, "GetTransactionsPerHour: failed to get data from repository")
	}
	return transactions, nil
}
