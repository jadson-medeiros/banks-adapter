package domain

type Transaction struct {
	Amount float64
	Type   int
	Text   string
}

const (
	CREDIT = 1
	DEBIT  = 2
)
