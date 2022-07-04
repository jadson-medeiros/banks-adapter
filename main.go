package main

import (
	"fmt"
	"time"

	"github.com/jadson-medeiros/banks-adapter/adapters"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

func main() {
	var bank1 domain.BanksProvider = &adapters.Bank1Adapter{}

	var bank2 domain.BanksProvider = &adapters.Bank2Adapter{}

	transactionsBank1 := bank1.GetTransactions(100, time.Now(), time.Now())
	transactionsBank2 := bank2.GetTransactions(100, time.Now(), time.Now())

	banksTransactions := append([]domain.Transaction{}, transactionsBank1...)

	banksTransactions = append(banksTransactions, transactionsBank2...)

	var banks []domain.BanksProvider
	banks = append(banks, bank1, bank2)

	PrintBalances(banks)
	PrintTransactions(banksTransactions)
}

func PrintBalances(banks []domain.BanksProvider) {
	for _, value := range banks {
		fmt.Printf("Balance: %v\n", value.GetBalance(100))
	}
}

func PrintTransactions(banks []domain.Transaction) {
	fmt.Printf("Transactions: %v\n", banks)
}
