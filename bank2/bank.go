package bank2

import "github.com/jadson-medeiros/banks-adapter/domain"

type Bank2Client struct{}

func (b *Bank2Client) GetBalance(accountID int64) domain.Balance {
	return domain.Balance{
		Balance:  512.5,
		Currency: "USD",
	}
}
