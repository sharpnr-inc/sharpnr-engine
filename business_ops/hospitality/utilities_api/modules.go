package hospitality_utilities

type HospitalityModulesType string

const (
	Restaurant             HospitalityModulesType = "RESTAURANT"
	Hotel                  HospitalityModulesType = "HOTEL"
	Bar                    HospitalityModulesType = "BAR"
	NightClub              HospitalityModulesType = "NIGHT_CLUB"
	Resort                 HospitalityModulesType = "RESORT"
	CloudKitchen_Takeaways HospitalityModulesType = "CLOUD_KITCHEN/TAKEAWAYS"
)

// Business module
type HospitalityModule struct {
	ID          string                 `json:"id"`
	ModuleName  HospitalityModulesType `json:"moduleName"`
	Description string                 `json:"description"`
	IsAvailable bool                   `json:"isAvailable"`
	Tags        []string               `json:"tags"`
}

// Infra module
type HospitalityInfraModule struct {
	ID          string   `json:"id"`
	ModuleName  string   `json:"moduleName"`
	Description string   `json:"description"`
	IsAvailable bool     `json:"isAvailable"`
	Tags        []string `json:"tags"`
}

var HospitalityModules = []HospitalityModule{
	{
		ID:          "HOTEL",
		ModuleName:  Hotel,
		Description: "Comprehensive hotel management module for handling reservations, room inventory, guest check-ins/outs, housekeeping, billing, and reporting.",
		IsAvailable: true,
		Tags:        []string{"accommodation", "reservations", "rooms", "billing"},
	},
	{
		ID:          "RESTAURANT",
		ModuleName:  Restaurant,
		Description: "End-to-end restaurant operations module for managing tables, menus, orders, kitchen display systems (KDS), billing, and online ordering integration.",
		IsAvailable: true,
		Tags:        []string{"food", "dining", "orders", "kds", "billing"},
	},
	{
		ID:          "NIGHT_CLUB",
		ModuleName:  NightClub,
		Description: "Entertainment and nightlife operations module covering guest entry, table and bottle service, membership management, event scheduling, and POS integration.",
		IsAvailable: false,
		Tags:        []string{"nightlife", "events", "membership", "pos"},
	},
	{
		ID:          "BAR",
		ModuleName:  Bar,
		Description: "Bar and lounge management system focusing on inventory tracking, liquor control, bartender sales monitoring, and daily revenue insights.",
		IsAvailable: true,
		Tags:        []string{"beverages", "liquor", "inventory", "sales"},
	},
	{
		ID:          "CLOUD_KITCHEN_TAKEAWAYS",
		ModuleName:  CloudKitchen_Takeaways,
		Description: "Specialized cloud kitchen and takeaway module designed for multi-brand virtual kitchens—supports order aggregation, delivery partner APIs, kitchen workflows, and real-time analytics.",
		IsAvailable: false,
		Tags:        []string{"delivery", "takeaway", "multi-brand", "analytics"},
	},
}

var HospitalityInfraModules = []HospitalityInfraModule{
	{
		ID:          "8789",
		ModuleName:  "Kitchen/Service Point",
		Description: "Centralized area for food preparation, order staging, and service coordination.",
		IsAvailable: true,
		Tags:        []string{"kitchen", "food-prep", "service"},
	},
	{
		ID:          "8756RQ",
		ModuleName:  "Bar Counter",
		Description: "Dedicated module for beverage preparation, mixology, and direct guest service.",
		IsAvailable: false,
		Tags:        []string{"bar", "beverages", "service"},
	},
	{
		ID:          "9910DH",
		ModuleName:  "Dining Hall",
		Description: "Manages physical floor plans, table assignments, and guest seating capacity.",
		IsAvailable: false,
		Tags:        []string{"seating", "floor-plan", "tables"},
	},
	{
		ID:          "4421ST",
		ModuleName:  "Inventory Store",
		Description: "Storage management for bulk supplies, dry goods, and non-perishable stock.",
		IsAvailable: false,
		Tags:        []string{"inventory", "storage", "supplies"},
	},
	{
		ID:          "1022FD",
		ModuleName:  "Reception/Front Desk",
		Description: "Primary guest interface for bookings, check-ins, and concierge services.",
		IsAvailable: false,
		Tags:        []string{"front-desk", "bookings", "concierge"},
	},
	{
		ID:          "5530HK",
		ModuleName:  "Housekeeping",
		Description: "Operational module for facility maintenance, room cleaning, and laundry services.",
		IsAvailable: false,
		Tags:        []string{"housekeeping", "cleaning", "maintenance"},
	},
}
