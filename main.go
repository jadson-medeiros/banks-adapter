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

	// bank2 := &Bank{
	// 	Method: &adapters.Bank2Adapter{},
	// }

	bank1.Method.GetTransactions(100, time.Now(), time.Now())

	fmt.Println("Starting")
}

func PrintBalances() {
	fmt.Println("Implement me to pull balance information from all available bank integrations and display them, one after the other.")
}

func PrintTransactions() {
	fmt.Println("Implement me to pull transactions from all available bank integrations and display them, one after the other.")
}
