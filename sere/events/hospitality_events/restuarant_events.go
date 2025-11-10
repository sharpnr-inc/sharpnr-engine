package hospitality_events

import "sharpnr-engine/sere"

const (

	//Event When Order is Creates
	EventOrderCreated sere.EventType = "hospitality_events.restaurant.order_created"
	// Event When Order is updated to Order in Progress
	EventOrderInProgress sere.EventType = "hospitality_events.restaurant.order_in_progress"
	EventOrderCompleted  sere.EventType = "hospitality_events.restaurant.order_completed"
	EventBillGenerated   sere.EventType = "hospitality_events.restaurant.bill_generated"
)
