package hospitalityconstants

/*
====================================================
 BAR ORDERS
====================================================
*/

type BarOrderStatus string

const (
	BarOrderPending   BarOrderStatus = "PENDING"
	BarOrderPreparing BarOrderStatus = "PREPARING"
	BarOrderServed    BarOrderStatus = "SERVED"
	BarOrderCancelled BarOrderStatus = "CANCELLED"
)

/*
====================================================
 BAR CHARGES / REVENUE
====================================================
*/

// BarChargeType represents bar income
type BarChargeType string

const (
	AlcoholSale BarChargeType = "ALCOHOL"
	SoftDrink   BarChargeType = "SOFT_DRINK"
	Snacks      BarChargeType = "SNACKS"
	ServiceFee  BarChargeType = "SERVICE_FEE"
)
