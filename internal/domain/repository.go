package domain

type TransactionsRepository interface {
	GetTransactionsPerHour(from, to int64) ([]TransactionCostPerHour, error)
}
