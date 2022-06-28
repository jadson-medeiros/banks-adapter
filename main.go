package main

import (
	"fmt"
	"time"

	"github.com/jadson-medeiros/banks-adapter/adapters"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

type Bank struct {
	Method domain.BanksMethod
}

func main() {
	bank1 := &Bank{
		Method: &adapters.Bank1Adapter{},
	}

	bank2 := &Bank{
		Method: &adapters.Bank2Adapter{},
	}

	transactionsBank1 := bank1.Method.GetTransactions(100, time.Now(), time.Now())
	transactionsBank2 := bank2.Method.GetTransactions(100, time.Now(), time.Now())

	banksTransactions := append([]domain.Transaction{}, transactionsBank1...)

	banksTransactions = append(banksTransactions, transactionsBank2...)

	var banks []Bank
	banks = append(banks, *bank1, *bank2)

	PrintBalances(banks)
	PrintTransactions(banksTransactions)
}

func PrintBalances(banks []Bank) {
	for _, value := range banks {
		fmt.Printf("Balance: %v\n", value.Method.GetBalance(100))
	}
}

func PrintTransactions(banks []domain.Transaction) {
	fmt.Printf("Transactions: %v\n", banks)
}
