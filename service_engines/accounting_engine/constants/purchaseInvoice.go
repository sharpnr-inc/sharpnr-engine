package accountingconstants

import "fmt"

/* ---------------------------------- <***> --------------------------------- */

type PurchaseInvoicePaymentStatusType string

const (
	PurchaseInvoicePaymentStatusUnpaid PurchaseInvoicePaymentStatusType = "UNPAID"
	PurchaseInvoicePaymentStatusPaid   PurchaseInvoicePaymentStatusType = "PAID"
)

/* ---------------------------------- <***> --------------------------------- */

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

func (s AccountingPurchaseInvoiceStatus) DisplayName() string {
	switch s {
	case PurchaseInvoiceStatusReturn:
		return "Item Returned / Debit Note"
	case PurchaseInvoiceStatusPartiallyPaid:
		return "Partially Paid"
	default:
		// Converts "PAID" to "Paid"
		return string(s[0]) + string(s[1:])
	}
}

func (s AccountingPurchaseInvoiceStatus) CanTransitionTo(next AccountingPurchaseInvoiceStatus) bool {
	switch s {
	case PurchaseInvoiceStatusDraft:
		// Draft can go anywhere
		return true
	case PurchaseInvoiceStatusCancelled, PurchaseInvoiceStatusPAID:
		// Final states usually cannot be changed
		return false
	case PurchaseInvoiceStatusPartiallyPaid:
		return next == PurchaseInvoiceStatusPAID || next == PurchaseInvoiceStatusReturn
	default:
		return false
	}
}

/* ---------------------------------- <***> --------------------------------- */
