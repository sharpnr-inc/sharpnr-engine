package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type RestaurantService interface {
	InitializeRestaurant()
	GetRestaurantDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO, error)
	UpdateRestaurantGeneralDetails(ctx context.Context, req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ, orgId uuid.UUID) (bool, error)

	CreateOrAddTable(ctx context.Context, orgid uuid.UUID, req hospitalitydtos.Restaurant_CreateOrAddTables_Request) error
	CreateFloor(ctx context.Context)
	GetRestaurantLayout(ctx context.Context, orgid uuid.UUID) (*hospitalitydtos.Restaurant_GetRestaurantLayout_Response, error)
}
