package kitchen_contracts

import (
	"time"

	"github.com/google/uuid"
)

type KOT_resp_dto struct {
	ID        uuid.UUID          `json:"id"`
	TicketNo  string             `json:"ticket_no"`
	Status    string             `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	Items     []KOTItem_resp_dto `json:"items"`
}

type KOTItem_resp_dto struct {
	OrderID      uuid.UUID `json:"order_id"`
	MenuItemID   uuid.UUID `json:"menu_item_id"`
	MenuItemName string
	VariantID    *uuid.UUID          `json:"variant_id,omitempty"`
	VaraintName  *string             `jsonr:"variant_name,omitempty"`
	Quantity     int                 `json:"quantity"`
	Remarks      string              `json:"remarks"`
	IsPrinted    bool                `json:"is_printed"`
	Addons       []KOTAddon_resp_dto `json:"addons"`
}

type KOTAddon_resp_dto struct {
	ID       uuid.UUID `json:"id"`
	AddOnID  uuid.UUID `json:"addon_id"`
	Quantity int       `json:"quantity"`
}

type KitchenConfigResponseDto struct {
	PrinterType string `json:"printer_type"`
	PrinterIP   string `json:"printer_ip"`
	PrinterPort int    `json:"printer_port"`
	AutoAccept  bool   `json:"auto_accept"`
	Throttling  struct {
		Enabled          bool `json:"enabled"`
		MaxOrdersPerSlot int  `json:"max_orders_per_slot"`
		SlotDurationMins int  `json:"slot_duration_mins"`
	} `json:"throttling"`

	OperatingHours []struct {
		DayOfWeek   int    `json:"day_of_week"`
		OpeningTime string `json:"opening_time"`
		ClosingTime string `json:"closing_time"`
	} `json:"operating_hours"`
}
