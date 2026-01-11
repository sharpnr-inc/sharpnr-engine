package hospitalityconstants

/*
====================================================
 TAX (SHARED)
====================================================
*/

// TaxType represents WHAT tax applies
type HospitalityTaxType string

const (
	RoomGST    HospitalityTaxType = "ROOM_GST"
	FoodGST    HospitalityTaxType = "FOOD_GST"
	ServiceGST HospitalityTaxType = "SERVICE_GST"
)

// TaxConditionType represents HOW tax is applied
type TaxConditionType string

const (
	TaxConditionPriceSlab TaxConditionType = "PRICE_SLAB"
	// TaxConditionStayNights TaxConditionType = "STAY_NIGHTS" // future
	// TaxConditionRoomType   TaxConditionType = "ROOM_TYPE"   // future
)

// TaxComponentType represents HOW tax is split
// Maps to AccountMapping.SourceCode (TAX)
type TaxComponentType string

const (
	TaxComponentCGST TaxComponentType = "CGST"
	TaxComponentSGST TaxComponentType = "SGST"
	TaxComponentIGST TaxComponentType = "IGST"
)

/*
====================================================
 PAYMENT (SHARED)
====================================================
*/

// PaymentMethod represents HOW money is received
// Maps to AccountMapping.SourceCode (PAYMENT)
type PaymentMethod string

const (
	PaymentCash         PaymentMethod = "CASH"
	PaymentUPI          PaymentMethod = "UPI"
	PaymentCard         PaymentMethod = "CARD"
	PaymentBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentRoomCharge   PaymentMethod = "ROOM_CHARGE"
)

// PaymentStatus represents payment lifecycle
type PaymentStatus string

const (
	PaymentSuccess PaymentStatus = "SUCCESS"
	PaymentPending PaymentStatus = "PENDING"
	PaymentFailed  PaymentStatus = "FAILED"
)

/*
====================================================
 BILLING (SHARED)
====================================================
*/

// BillStatus represents billing lifecycle
type BillStatus string

const (
	BillDraft     BillStatus = "DRAFT"
	BillFinalized BillStatus = "FINALIZED"
	BillCancelled BillStatus = "CANCELLED"
)

type SourceType string

const (
	// -------------------------
	// Revenue / Cost generation
	// -------------------------
	SourceCharge  SourceType = "CHARGE"  // room rent, food, service, minibar
	SourceTax     SourceType = "TAX"     // CGST, SGST, IGST
	SourcePayment SourceType = "PAYMENT" // cash, card, UPI, bank

	// -------------------------
	// System generated
	// -------------------------
	SourceSystem SourceType = "SYSTEM" // receivable, payable, rounding

	// -------------------------
	// Adjustments
	// -------------------------
	SourceAdjustment SourceType = "ADJUSTMENT" // discounts, write-offs
	SourceRefund     SourceType = "REFUND"     // customer refund
)
