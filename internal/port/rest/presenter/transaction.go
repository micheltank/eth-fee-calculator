package presenter

import (
	"github.com/micheltank/eth-fee-calculator/internal/domain"
)

type TransactionCostsPerHourParams struct {
	From int64 `form:"from" validate:"required"`
	To   int64 `form:"to" validate:"required"`
	Page int   `form:"page"`
}

type TransactionCostPerHourResponse struct {
	Time  int64   `json:"t"`
	Value float64 `json:"v"`
}

func NewTransactionCostPerHourResponse(transaction domain.TransactionCostPerHour) TransactionCostPerHourResponse {
	return TransactionCostPerHourResponse{
		Time:  transaction.Hour.Unix(),
		Value: transaction.FeeAmount,
	}
}

func NewTransactionCostsPerHourResponse(transactions []domain.TransactionCostPerHour) (response []TransactionCostPerHourResponse) {
	for _, transaction := range transactions {
		response = append(response, NewTransactionCostPerHourResponse(transaction))
	}
	return response
}
