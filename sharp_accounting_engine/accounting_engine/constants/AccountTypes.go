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

// AccountGroupTypeResponse encapsulates account group field requirements
type AccountGroupTypeResponse struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// Label returns the standard name of the account group
func (t AccountGroupType) Label() string {
	switch t {
	case AccountAsset:
		return "Asset"
	case AccountLiability:
		return "Liability"
	case AccountEquity:
		return "Equity"
	case AccountIncome:
		return "Income"
	case AccountExpense:
		return "Expense"
	default:
		return string(t)
	}
}

// Description explains the accounting purpose for frontend tooltips or context hints
func (t AccountGroupType) Description() string {
	switch t {
	case AccountAsset:
		return "Resources owned or controlled by the business that are expected to provide future economic benefits."
	case AccountLiability:
		return "Obligations or debts owed by the business to external parties, payable in the future."
	case AccountEquity:
		return "The residual interest in the assets of the business after deducting all its liabilities, representing owner's claim."
	case AccountIncome:
		return "Money received or earned by the business from its operational activities, which increases equity."
	case AccountExpense:
		return "Costs incurred or resources consumed by the business to generate revenue, which decreases equity."
	default:
		return ""
	}
}
