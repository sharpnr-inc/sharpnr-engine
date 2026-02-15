package restaurant_constants

type RestaurantDiningSessionStatusType string

const (
	Opened RestaurantDiningSessionStatusType = "OPEN"
	Closed RestaurantDiningSessionStatusType = "CLOSED"
)

type RestaurantTableStatusType string

const (
	Available     RestaurantTableStatusType = "AVAILABLE"
	Occupied      RestaurantTableStatusType = "OCCUPIED"
	Reserved      RestaurantTableStatusType = "RESERVED"
	Blocked       RestaurantTableStatusType = "BLOCKED"
	UnderCleaning RestaurantTableStatusType = "UNDER_CLEANING"
)

type RestaurantRunningStatusType string

const (
	Active    RestaurantRunningStatusType = "ACTIVE"
	NotActive RestaurantRunningStatusType = "NOT_ACTIVE"
)

type RestaurantConfigurationType string

const (
	Configured    RestaurantConfigurationType = "CONFIGURED"
	NotConfigured RestaurantConfigurationType = "NOT_CONFIGURED"
)

type RestaurantOrderType string

const (
	DineInOrder  RestaurantOrderType = "DINE_IN"
	TakewayOrder RestaurantOrderType = "TAKEAWAY"
)
