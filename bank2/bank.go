package bank2

type Balance struct {
	Balance  float64
	Currency string
}

func GetBalance(accountID int64) Balance {
	return Balance{
		Balance:  512.5,
		Currency: "USD",
	}
}
