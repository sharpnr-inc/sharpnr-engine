package accountingconstants

type AccountGroupType string

const (
	AccountAsset     AccountGroupType = "ASSET"
	AccountLiability AccountGroupType = "LIABILITY"
	AccountEquity    AccountGroupType = "EQUITY"
	AccountIncome    AccountGroupType = "INCOME"
	AccountExpense   AccountGroupType = "EXPENSE"
)

// AllAccountGroups provides a slice for easy iteration
var AllAccountGroups = []AccountGroupType{
	AccountAsset,
	AccountLiability,
	AccountEquity,
	AccountIncome,
	AccountExpense,
}
