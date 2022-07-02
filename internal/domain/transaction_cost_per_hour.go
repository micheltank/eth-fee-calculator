package domain

import "time"

type TransactionCostPerHour struct {
	Hour      time.Time
	FeeAmount float64
}
