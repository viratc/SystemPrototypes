package models

type Expense struct {
	UserBalances map[string]int
	Title        string
	ImageUrl     string
	Description  string
	GroupId      string
}

func NewExpense(balances map[string]int, title string, imageUrl string, description string, groupId string) *Expense {
	return &Expense{UserBalances: balances, Title: title, ImageUrl: imageUrl, Description: description, GroupId: groupId}
}

func (e *Expense) GetUserBalances() map[string]int {
	return e.UserBalances
}

func (e *Expense) GetGroupId() string {
	return e.GroupId
}
