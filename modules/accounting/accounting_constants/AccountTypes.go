package accountingconstants

type AccountType string

const (
	AccountAsset     AccountType = "ASSET"
	AccountLiability AccountType = "LIABILITY"
	AccountEquity    AccountType = "EQUITY"
	AccountIncome    AccountType = "INCOME"
	AccountExpense   AccountType = "EXPENSE"
)
