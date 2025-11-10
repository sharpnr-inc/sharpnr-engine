package hospitality_events

import "sharpnr-engine/sere"

// These Events are related to Kitchen
const (
	KitchenEventOrderReceived sere.EventType = "hospitality.kitchen.order_received"
	KitchenEventOrderInProg   sere.EventType = "hospitality.kitchen.order_in_progress"
	KitchenEventOrderServed   sere.EventType = "hospitality.kitchen.order_served"
)
