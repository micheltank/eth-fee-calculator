package repository

import (
	"database/sql"

	"github.com/micheltank/eth-fee-calculator/internal/domain"
)

type TransactionPostgreSql struct {
	db *sql.DB
}

func NewTransactionPostgreSql(db *sql.DB) *TransactionPostgreSql {
	return &TransactionPostgreSql{
		db: db,
	}
}

func (r *TransactionPostgreSql) GetTransactionsPerHour(from, to int64, page int) ([]domain.TransactionCostPerHour, error) {
	limit, offset := getLimitAndOffset(page)

	rows, err := r.db.Query(`SELECT
							date_trunc('hour', t.block_time) as hour,
							round(sum((t.gas_used * t.gas_price)/power(10, 18))::numeric, 2) as gas_cost
						FROM transactions t
						WHERE t.from != '0x0000000000000000000000000000000000000000'
						  AND t.to != '0x0000000000000000000000000000000000000000'
						  AND EXTRACT(EPOCH FROM (t.block_time AT TIME ZONE 'UTC')) BETWEEN $1 AND $2
						GROUP BY hour
						ORDER BY hour
						LIMIT $3 OFFSET $4`, from, to, limit, offset)
	if err != nil {
		return []domain.TransactionCostPerHour{}, err
	}
	defer rows.Close()
	var transactions []domain.TransactionCostPerHour
	for rows.Next() {
		var transaction domain.TransactionCostPerHour
		err = rows.Scan(&transaction.Hour,
			&transaction.FeeAmount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func getLimitAndOffset(page int) (int, int) {
	limit := 12
	if page <= 0 {
		page = 1
	}
	offset := limit * (page - 1)
	return limit, offset
}
