package restaurantservices

import "context"

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
func (s *restaurantService) CreateTableBookingSession(ctx context.Context, request any) error {
	panic("unimplemented")
}

func (s *restaurantService) AddItemInBookingSession(ctx context.Context, request any) error {
	panic("unimplemented")
}
