package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

/****************************************
 * Service Interface for the Restuarant *
 *                                      *
****************************************/
// RestaurantService defines all restaurant-related operations
// grouped cleanly by domain (Restaurant → Layout → Session → Orders)
type RestaurantService interface {

	// -------------------------------------------------------------------------
	// 🏪 RESTAURANT DETAILS
	// -------------------------------------------------------------------------

	// Returns organization-specific restaurant details
	GetRestaurantDetails(
		ctx context.Context,
		orgId uuid.UUID,
	) (*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO, error)

	// Updates the general details like name, GST info, address etc.
	UpdateRestaurantGeneralDetails(
		ctx context.Context,
		req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ,
		orgId uuid.UUID,
	) (bool, error)

	// -------------------------------------------------------------------------
	// 🗺️ RESTAURANT LAYOUT (Floors, Tables)
	// -------------------------------------------------------------------------

	// Creates or adds one or multiple tables in a specific floor
	CreateOrAddTable(
		ctx context.Context,
		orgid uuid.UUID,
		req hospitalitydtos.Restaurant_CreateOrAddTables_Request,
	) error

	// Creates a new floor inside the restaurant
	CreateFloor(ctx context.Context)

	// Returns complete layout → floors + tables + statuses
	GetRestaurantLayout(
		ctx context.Context,
		orgid uuid.UUID,
	) (*hospitalitydtos.Restaurant_GetRestaurantLayout_Response, error)

	// -------------------------------------------------------------------------
	// 🕒 TABLE BOOKING SESSION
	// -------------------------------------------------------------------------

	// Creates a new table booking session (opens a table)
	CreateTableBookingSession(
		ctx context.Context,
		orgId uuid.UUID,
		request hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST,
	) error

	// Gets active session for a given table (returns nil if no session)
	GetResTableBookingSessionByTableId(
		ctx context.Context,
		orgId, tableId uuid.UUID,
	) (hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response, error)

	// -------------------------------------------------------------------------
	// 🍽️ ORDERS WITHIN A SESSION
	// -------------------------------------------------------------------------

	// Adds a new item/order to an existing table session
	AddItemInBookingSession(
		ctx context.Context,
		orgId, sessionId uuid.UUID,
		request hospitalitydtos.Hospitality_Restaurant_AddOrderItemINTheSession_Request,
	) error
}
