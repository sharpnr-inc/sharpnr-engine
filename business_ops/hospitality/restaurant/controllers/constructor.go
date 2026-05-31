package restaurantcontrollers

import restaurantservices "sharpnr-engine/business_ops/hospitality/restaurant/services"

// RestaurantController handles all restaurant-related HTTP requests.
//
// This controller acts as the bridge between:
//   - API routes / handlers
//   - Restaurant business logic services
//
// Responsibilities:
//   - Receive incoming restaurant requests
//   - Validate/request process inputs
//   - Call restaurant service methods
//   - Return responses back to clients
type RestaurantController struct {
	// resSvc contains the restaurant business logic layer.
	//
	// The controller depends on the service interface
	// instead of a concrete implementation to support:
	//   - Dependency Injection
	//   - Better testing/mocking
	//   - Cleaner architecture
	resSvc restaurantservices.RestaurantServiceInterface
}

// NewRestaurantController creates and returns a new RestaurantController instance.
//
// Parameters:
//   - resSvc: Restaurant service implementation that contains
//     restaurant business logic operations.
//
// Returns:
//   - Pointer to initialized RestaurantController
//
// This constructor is typically used during application bootstrap
// when wiring controllers and services together.
func NewRestaurantController(
	resSvc restaurantservices.RestaurantServiceInterface,
) *RestaurantController {
	return &RestaurantController{
		resSvc: resSvc,
	}
}
