package adapters

import (
	"time"

	"github.com/jadson-medeiros/banks-adapter/bank1"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

type Bank1Adapter struct {
	Bank bank1.Bank1Client
}

func (b *Bank1Adapter) GetBalance(accountID int64) float64 {
	return b.Bank.GetAccountBalance(accountID)
}

func (b *Bank1Adapter) GetCurrency(accountID int64) string {
	return b.Bank.GetAccountCurrency(accountID)
}

func (b *Bank1Adapter) GetTransactions(account int64, fromDate time.Time, toDate time.Time) []domain.Transaction {
	checkDeposit := domain.Transaction{Amount: 100, Type: domain.CREDIT, Text: "Check deposit"}
	debidCardPurchase := domain.Transaction{Amount: 25.5, Type: domain.DEBIT, Text: "Debit card purchase"}
	rentPayment := domain.Transaction{Amount: 225, Type: domain.DEBIT, Text: "Rent payment"}

	return []domain.Transaction{checkDeposit, debidCardPurchase, rentPayment}
}
