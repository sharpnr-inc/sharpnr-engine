package accounting_capabilities

type AccountingModuleCapabilities string

const (
	AddTransactions AccountingModuleCapabilities = "ACCOUNTING:ADD_TRANSACTIONS"
	ViewLedger      AccountingModuleCapabilities = "ACCOUNTING:VIEW_LEDGER"
	ViewReports     AccountingModuleCapabilities = "ACCOUNTING:VIEW_REPORTS"
)
