package main

import (
	"fmt"
	"time"

	"github.com/jadson-medeiros/banks-adapter/adapters"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

type BanksMethod interface {
	GetTransactions(account int64, fromDate time.Time, toDate time.Time) []domain.Transaction
	GetBalance(accountId int64) float64
	GetCurrency(accountId int64) string
}

type Bank struct {
	Method BanksMethod
}

func main() {

	bank1 := &Bank{
		Method: &adapters.Bank1Adapter{},
	}

	bank2 := &Bank{
		Method: &adapters.Bank2Adapter{},
	}

	bank1.Method.GetTransactions(100, time.Now(), time.Now())

	fmt.Println("Starting")
}
