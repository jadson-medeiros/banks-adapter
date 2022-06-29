package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/jadson-medeiros/banks-adapter/adapters"
	"github.com/jadson-medeiros/banks-adapter/domain"
)

func CreateBank1Mock() *Bank {
	return &Bank{
		Method: &adapters.Bank1Adapter{},
	}
}

func CreateBank2Mock() *Bank {
	return &Bank{
		Method: &adapters.Bank2Adapter{},
	}
}

func GetTransactionsBank1Mock() []domain.Transaction {
	checkDeposit := domain.Transaction{Amount: 100, Type: domain.CREDIT, Text: "Check deposit"}
	debidCardPurchase := domain.Transaction{Amount: 25.5, Type: domain.DEBIT, Text: "Debit card purchase"}
	rentPayment := domain.Transaction{Amount: 225, Type: domain.DEBIT, Text: "Rent payment"}

	return []domain.Transaction{checkDeposit, debidCardPurchase, rentPayment}
}

func GetTransactionsBank2Mock() []domain.Transaction {
	purchase := domain.Transaction{Amount: 125, Type: domain.DEBIT, Text: "Amazon.com"}
	carInsurance := domain.Transaction{Amount: 500, Type: domain.DEBIT, Text: "Car insurance"}
	salary := domain.Transaction{Amount: 800, Type: domain.CREDIT, Text: "Salary"}

	return []domain.Transaction{purchase, carInsurance, salary}
}

func TestGetTransactionsBank1(t *testing.T) {
	bank := CreateBank1Mock()

	result := bank.Method.GetTransactions(100, time.Now(), time.Now())

	expected := GetTransactionsBank1Mock()

	if result == nil {
		t.Errorf("result is nill")
		return
	}

	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("result '%s', expected '%s'", fmt.Sprint(result), fmt.Sprint(expected))
	}
}

func TestGetTransactionsBank2(t *testing.T) {
	bank := CreateBank2Mock()

	result := bank.Method.GetTransactions(100, time.Now(), time.Now())

	expected := GetTransactionsBank2Mock()

	if result == nil {
		t.Errorf("result is nill")
		return
	}

	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("result '%s', expected '%s'", fmt.Sprint(result), fmt.Sprint(expected))
	}
}
