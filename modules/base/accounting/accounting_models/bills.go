package accountingmodels

type AccountingBills struct {
	ID          string
	CustomerId  string
	Description string
	InvoiceDate string
	Amount      string
}

func (AccountingBills) TableName() string {
	return "accounting_bill"
}
