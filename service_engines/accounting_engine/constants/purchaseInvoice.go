package accountingconstants

import "fmt"

type AccountingPurchaseInvoiceStatus string

const (
	PurchaseInvoiceStatusDraft         AccountingPurchaseInvoiceStatus = "DRAFT"
	PurchaseInvoiceStatusPAID          AccountingPurchaseInvoiceStatus = "PAID"
	PurchaseInvoiceStatusPartiallyPaid AccountingPurchaseInvoiceStatus = "PARTIALLY_PAID"
	PurchaseInvoiceStatusReturn        AccountingPurchaseInvoiceStatus = "ITEM_RETURN/DEBIT_NOTE"
	PurchaseInvoiceStatusCancelled     AccountingPurchaseInvoiceStatus = "CANCELLED"
)

func (s AccountingPurchaseInvoiceStatus) VerifyPurchaseTrxType() (string, error) {
	switch s {
	case PurchaseInvoiceStatusDraft,
		PurchaseInvoiceStatusPAID,
		PurchaseInvoiceStatusPartiallyPaid,
		PurchaseInvoiceStatusReturn,
		PurchaseInvoiceStatusCancelled:
		return string(s), nil

	default:
		return "", fmt.Errorf("invalid purchase invoice status: %s", s)
	}
}
