package restaurant_constants

type RestaurantDiningSessionStatusType string

const (
	DineInSessionOpened RestaurantDiningSessionStatusType = "OPEN"
	DineInSessionClosed RestaurantDiningSessionStatusType = "CLOSED"
)

type RestaurantTableStatusType string

const (
	TableStatusAvailable     RestaurantTableStatusType = "AVAILABLE"
	TableStatusOccupied      RestaurantTableStatusType = "OCCUPIED"
	TableStatusReserved      RestaurantTableStatusType = "RESERVED"
	TableStatusBlocked       RestaurantTableStatusType = "BLOCKED"
	TableStatusUnderCleaning RestaurantTableStatusType = "UNDER_CLEANING"
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
