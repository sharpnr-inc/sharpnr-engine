package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
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
	ctx context.Context,
	request hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST) error {

	/*****************
	 * Function Body *
	 *****************/

	// TODO : ORGID (Fetched via Middleware)--> RestaurantID --> can be fetched on the basis of the OrgID
	// *1 Need Repo function for the Getting the Table By the ID
	// *2 Then we can matched it with the RestaurantID and Varify the Status
	// Then Proceeds to Creating the Session Order and Initial Order Slice
	// and since we have the KitchenRepo Access we can find the KitchenBy ID of OrgID
	// and Create the KitchenTicket with the KitchenTicketItems

	// Which will send the Data to Kitchen Via Websocket with the SERE(Sharpnr Event Relay Engine)
	return nil
}

func (s *restaurantService) AddItemInBookingSession(ctx context.Context, request any) error {
	panic("unimplemented")
}
