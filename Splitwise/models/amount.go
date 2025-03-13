package models

type Amount struct {
	Currency Currency
	Amount   int64
}

func NewAmount(currency Currency, amount int64) *Amount {
	return &Amount{Currency: currency, Amount: amount}
}

func (a *Amount) GetCurrency() Currency {
	return a.Currency
}

func (a *Amount) GetAmount() int64 {
	return a.Amount
}

func (a *Amount) Add(amt int64) *Amount {
	return &Amount{Currency: a.Currency, Amount: a.Amount + amt}
}
