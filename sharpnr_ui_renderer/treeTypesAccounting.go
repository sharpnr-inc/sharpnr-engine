package sharpnruirenderer

// ----------------------------Accounting Root-----------------------------------
type SharpAccountingTree struct {
	HasAccess             bool                  `json:"has_access"`
	SharpAccountingModule SharpAccountingModule `json:"sub_module_section"`
}

// ----------------------------Accounting Modules-----------------------------------
type SharpAccountingModule struct {
	Transactions        TransactionsModule        `json:"transactions"`
	Expenses            ExpensesModule            `json:"expenses"`
	Invoices            InvoicesModule            `json:"invoices"`
	Ledger              LedgerModule              `json:"ledger"`
	Customers           CustomersModule           `json:"customers"`
	SuppliersAndVendors SuppliersAndVendorsModule `json:"suppliersandvendors"`
	Reports             ReportsModule             `json:"reports"`
}

// ----------------------------Transactions-----------------------------------
type TransactionsModule struct {
	HasAccess             bool                  `json:"has_access"`
	TransactionSubSection TransactionSubSection `json:"sub_section_access"`
}

type TransactionSubSection struct {
	// define later (e.g. journal_entries, transfers)
}

// ----------------------------Expenses-----------------------------------
type ExpensesModule struct {
	HasAccess         bool              `json:"has_access"`
	ExpenseSubSection ExpenseSubSection `json:"sub_section_access"`
}

type ExpenseSubSection struct{}

// ----------------------------Invoices-----------------------------------
type InvoicesModule struct {
	HasAccess         bool              `json:"has_access"`
	InvoiceSubSection InvoiceSubSection `json:"sub_section_access"`
}

type InvoiceSubSection struct{}

// ----------------------------Ledger-----------------------------------
type LedgerModule struct {
	HasAccess        bool             `json:"has_access"`
	LedgerSubSection LedgerSubSection `json:"sub_section_access"`
}

type LedgerSubSection struct{}

// ----------------------------Customers-----------------------------------
type CustomersModule struct {
	HasAccess          bool               `json:"has_access"`
	CustomerSubSection CustomerSubSection `json:"sub_section_access"`
}

type CustomerSubSection struct{}

// ----------------------------Suppliers & Vendors-----------------------------------
type SuppliersAndVendorsModule struct {
	HasAccess                bool                     `json:"has_access"`
	SupplierVendorSubSection SupplierVendorSubSection `json:"sub_section_access"`
}

type SupplierVendorSubSection struct{}

// ----------------------------Reports-----------------------------------
type ReportsModule struct {
	HasAccess        bool             `json:"has_access"`
	ReportSubSection ReportSubSection `json:"sub_section_access"`
}

type ReportSubSection struct{}
