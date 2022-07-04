package domain

import "time"

type BanksProvider interface {
	GetTransactions(account int64, fromDate time.Time, toDate time.Time) []Transaction
	GetBalance(accountId int64) float64
	GetCurrency(accountId int64) string
}
