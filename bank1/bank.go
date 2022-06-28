package bank1

type Bank1Client struct{}

func (b Bank1Client) GetAccountBalance(accountID int64) float64 {
	return 215.5
}

func (b Bank1Client) GetAccountCurrency(accountID int64) string {
	return "USD"
}
