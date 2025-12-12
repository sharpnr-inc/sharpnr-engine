package restaurantservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
)

/************************************************
 * This is the Service belongs to order for the *
 * All the Order belongs to Restaurant But      *
 ************************************************/

/*
This is Service Function Creating the Session of the Restaurant
Session is how many Orders and how long it took for a customer to rotate over a table.
This must be closed after the payment; till then it will remain open.
"request" is the body of the request used to create a session.
*/
func (s *restaurantService) CreateTableBookingSession(
	ctx context.Context, orgId uuid.UUID,
	request hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST) error {

	/*****************
	 * Function Body *
	 *****************/

	// TODO : ORGID (Fetched via Middleware)--> RestaurantID --> can be fetched on the basis of the OrgID
	// *1 Need Repo function for the Getting the Table By the ID
	// *2 Then we can matched it with the RestaurantID and Verify the Status
	// Then Proceeds to Creating the Session Order and Initial Order Slice
	// and since we have the KitchenRepo Access we can find the KitchenBy ID of OrgID
	// and Create the KitchenTicket with the KitchenTicketItems

	// Which will send the Data to Kitchen Via Websocket with the SERE(Sharpnr Event Relay Engine)

	restaurant, err := s.restaurantRepo.GetRestaurantByOrgId(ctx, orgId)
	if err != nil {
		return nil
	}

	table, err := s.restaurantRepo.GetTableById(ctx, request.TableId)
	if err != nil {
		return err
	}

	if table.RestaurantID != restaurant.ID {
		return errors.New("table does not belong to this restaurant")
	}

	session := &hospitalitymodels.RestaurantTableBookingSession{
		ID:                            uuid.New(),
		RestaurantTableID:             restaurant.ID,
		RestaurantTable:               *table,
		HospitalityCustomerAndGuestID: nil,
		CustomerAndGuest:              nil,
		FromTime:                      time.Now(),
		CreatedAt:                     time.Now(),
		UpdatedAt:                     time.Now(),
	}

	if err := s.restaurantRepo.CreateRes_TableSession(ctx, session); err != nil {
		return err
	}

	// Set table as OCCUPIED
	err = s.restaurantRepo.UpdateTable(
		ctx,
		table.ID,
		map[string]any{
			"status": "OCCUPIED",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *restaurantService) AddItemInBookingSession(
	ctx context.Context,
	orgId, sessionId uuid.UUID,
	request hospitalitydtos.Hospitality_Restaurant_AddOrderItemINTheSession_Request) error {
	return nil
}

func (s *restaurantService) GetResTableBookingSessionByTableId(
	ctx context.Context,
	orgId, tableId uuid.UUID,
) (hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response, error) {

	session, err := s.restaurantRepo.GetBookingSessionByTableId(ctx, tableId)
	if err != nil {
		return hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response{}, err
	}

	// session may be nil → no active session yet
	return hospitalitydtos.Hospitality_Restaurant_GetBookingSession_Response{
		Session: session,
	}, nil
}

func (s *restaurantService) GetOrderStatus(ctx context.Context, orgId, kitchenTicketId uuid.UUID) (interface{}, error) {
	return nil, nil
}
