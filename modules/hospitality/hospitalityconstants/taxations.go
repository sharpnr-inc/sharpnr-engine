package hospitalityconstants

type TaxConditionType string

const (
	TaxConditionPriceSlab TaxConditionType = "PRICE_SLAB"
	// TaxConditionStayNights TaxConditionType = "STAY_NIGHTS" // future
	// TaxConditionRoomType   TaxConditionType = "ROOM_TYPE"   // future
)

type TaxType string

const (
	TaxTypeRoomGST TaxType = "ROOM_GST"
	TaxTypeFoodGST TaxType = "FOOD_GST"
	// TaxTypeService TaxType = "SERVICE_CHARGE" // future
)

type TaxComponentType string

const (
	TaxComponentCGST TaxComponentType = "CGST"
	TaxComponentSGST TaxComponentType = "SGST"
	TaxComponentIGST TaxComponentType = "IGST"
)
