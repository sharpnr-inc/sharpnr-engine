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

type RestaurantOrderType string

const (
	DineInOrder  RestaurantOrderType = "DINE_IN"
	TakewayOrder RestaurantOrderType = "TAKEAWAY"
)

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                   Restaurant Configuration Related Enums                   */
/* -------------------------------------------------------------------------- */

// Restaurant Configuration
type RestaurantConfigurationStatusType string

const (
	RestaurantStatusCreated             RestaurantConfigurationStatusType = "CREATED"
	RestaurantStatusGeneralDetailsAdded RestaurantConfigurationStatusType = "GENERAL_DETAILS_ADDED"
	RestaurantStatusTableSetup          RestaurantConfigurationStatusType = "TABLE_SETUP"
	RestaurantStatusCompleted           RestaurantConfigurationStatusType = "COMPLETED"
)

/* ---------------------------------- <***> --------------------------------- */
