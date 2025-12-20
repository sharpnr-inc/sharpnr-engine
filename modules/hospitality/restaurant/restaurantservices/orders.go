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

	// 1️⃣ Get Restaurant
	restaurant, err := s.restaurantRepo.GetRestaurantByOrgId(ctx, orgId)
	if err != nil {
		return err
	}

	// 2️⃣ Get Session
	session, err := s.restaurantRepo.GetBookingSessionById(ctx, sessionId)
	if err != nil {
		return err
	}

	// 3️⃣ Create Kitchen Ticket
	orderTicket := &hospitalitymodels.KitchenTicket{
		ID:                              uuid.New(),
		OrganizationID:                  orgId,
		Organization:                    *restaurant.Organization,
		KitchenID:                       *restaurant.KitchenID,
		Kitchen:                         *restaurant.Kitchen,
		Status:                          "PENDING",
		TicketNo:                        "0001",
		RestaurantTableBookingSessionID: &session.ID,
		CreatedAt:                       time.Now(),
		UpdatedAt:                       time.Now(),
	}

	kot, err := s.kitchenRepo.CreateKitchenTicketDBTx(ctx, orderTicket)
	if err != nil {
		return err
	}

	// 4️⃣ Create All Items + Addons
	for _, items := range request.Order {

		// 4.1 → Insert item
		item := &hospitalitymodels.KitchenTicketItem{
			ID:              uuid.New(),
			KitchenTicketID: kot.ID,
			MenuItemID:      items.MenuItemId,
			VariantId:       items.VariantId,
			Quantity:        items.Quantity,
			Remarks:         items.Notes,
			OrderType:       "TableOrder",
			IsPrinted:       false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		item, err := s.kitchenRepo.CreateKitchenTicketItemDBTx(ctx, item)
		if err != nil {
			return err
		}

		// 4.2 → Add addon rows
		if items.AddonOnId != nil {
			for _, addon := range *items.AddonOnId {
				addonModel := &hospitalitymodels.KitchenTicketItemAddon{
					ID:                  uuid.New(),
					KitchenTicketItemID: item.ID,
					AddOnID:             addon.AddOnId,
					Quantity:            addon.Quantity,
					CreatedAt:           time.Now(),
				}

				if _, err := s.kitchenRepo.CreateKitchenTicketItemAddonDBTx(ctx, addonModel); err != nil {
					return err
				}
			}
		}
	}

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
