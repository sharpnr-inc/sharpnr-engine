package kitchen_contracts

import "github.com/google/uuid"

type KOTDisplayView struct {
	TicketID    uuid.UUID
	TicketNo    string
	TableSource string
	Items       []KOTItemView
}

type KOTItemView struct {
	Name     string
	Variant  string   // From VariantID
	Addons   []string // From KitchenTicketItemAddon
	Quantity int
	Status   string
	Remarks  string
}
