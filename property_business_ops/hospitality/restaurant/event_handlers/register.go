package restaurant_event_handlers

import (
	"context"
	"fmt"
	restaurant_events "sharpnr-engine/property_business_ops/hospitality/restaurant/events"
	restaurantservices "sharpnr-engine/property_business_ops/hospitality/restaurant/services"
	"sharpnr-engine/sero"
)

func Register(seroEngine *sero.SEROEngine, service restaurantservices.RestaurantServiceInterface) {

	// Subscribe to the OrderStatusUpdated event
	seroEngine.Subscribe(restaurant_events.OrderStatusUpdated, func(ctx context.Context, payload interface{}) error {

		var data restaurant_events.OrderStatusUpdatePayload

		// Use your MapToStruct helper to safely get the data
		if err := sero.MapToStruct(payload, &data); err != nil {
			return fmt.Errorf("restaurant_handler: failed to map payload: %w", err)
		}

		// Update the restaurant's order record
		return service.UpdateOrderStatus(ctx, data.OrderID, data.Status)

	}, false) // 'false' for Sync if you want it in the same transaction, 'true' for Async
}
