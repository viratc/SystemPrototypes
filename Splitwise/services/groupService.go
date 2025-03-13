package services

type GroupService struct {
}

// In memory map of groupIds
groups := map[string]Group

func GetGroupPaymentGraph(groupId string, userId string) {

	// Check if user belongs to this group

	groupExpenses := GetGroupExpenses()
	resultExpense := SumExpenses(groupExpenses)

	return GetPaymentGraph(resultExpense)
}

func SumExpenses() {

}

func GetBalances() {
	// Check if user belongs to the group
	groupExpenses := GetGroupExpenses()
	resultExpense := SumExpenses(groupExpenses)

	return GetPaymentGraph(resultExpense)
}
