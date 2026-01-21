package accountingconstants

type AccountGroupType string

const (
	AccountAsset     AccountGroupType = "ASSET"
	AccountLiability AccountGroupType = "LIABILITY"
	AccountEquity    AccountGroupType = "EQUITY"
	AccountIncome    AccountGroupType = "INCOME"
	AccountExpense   AccountGroupType = "EXPENSE"
)
