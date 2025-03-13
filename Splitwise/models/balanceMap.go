package models

import "fmt"

type BalanceMap struct {
	Balances map[string]Amount
}

func NewBalanceMap() *BalanceMap {
	return &BalanceMap{Balances: make(map[string]Amount)}
}

func NewBalanceMapWithBalances(balance map[string]Amount) *BalanceMap {
	return &BalanceMap{Balances: balance}
}

func (b *BalanceMap) GetBalance() map[string]Amount {
	return b.Balances
}

func (b *BalanceMap) String() string {
	return fmt.Sprintf("BalanceMap{Balances=%v}", b.Balances)
}
