package adapters

import (
	"time"

	"github.com/jadson-medeiros/banks-adapter/bank2"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

type Bank2Adapter struct {
	Bank bank2.Bank2Client
}

func (b *Bank2Adapter) GetBalance(accountID int64) float64 {
	balance := b.Bank.GetBalance(accountID)
	return balance.Balance
}

func (b *Bank2Adapter) GetCurrency(accountID int64) string {
	balance := b.Bank.GetBalance(accountID)
	return balance.Currency
}

func (b *Bank2Adapter) GetTransactions(account int64, fromDate time.Time, toDate time.Time) []domain.Transaction {
	purchase := domain.Transaction{Amount: 125, Type: domain.DEBIT, Text: "Amazon.com"}
	carInsurance := domain.Transaction{Amount: 500, Type: domain.DEBIT, Text: "Car insurance"}
	salary := domain.Transaction{Amount: 800, Type: domain.CREDIT, Text: "Salary"}

	return []domain.Transaction{purchase, carInsurance, salary}
}
