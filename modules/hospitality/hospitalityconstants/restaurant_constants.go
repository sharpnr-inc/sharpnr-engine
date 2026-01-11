package hospitalityconstants

/*
====================================================
 RESTAURANT ORDERS
====================================================
*/

type RestaurantOrderStatus string

const (
	RestaurantOrderPending   RestaurantOrderStatus = "PENDING"
	RestaurantOrderPreparing RestaurantOrderStatus = "PREPARING"
	RestaurantOrderServed    RestaurantOrderStatus = "SERVED"
	RestaurantOrderCancelled RestaurantOrderStatus = "CANCELLED"
)

/*
====================================================
 RESTAURANT CHARGES / REVENUE
====================================================
*/

// RestaurantChargeType represents restaurant income
type RestaurantChargeType string

const (
	Restaurant_Food            RestaurantChargeType = "FOOD"
	Restaurant_PackagingCharge RestaurantChargeType = "PACKAGING"
	Restaurant_ServiceCharge   RestaurantChargeType = "SERVICE_CHARGE"
)

/*
====================================================
 TABLE / SESSION
====================================================
*/

type TableSessionStatus string

const (
	RestaurantTableSessionOpen   TableSessionStatus = "OPEN"
	RestaurantTableSessionClosed TableSessionStatus = "CLOSED"
)
