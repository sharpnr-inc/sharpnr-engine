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

type HospitalityDomainModule struct {
	ID          string
	ModuleName  HospitalityModulesType
	Description string
	IsAvailable bool
}

var HospitalityModules = []HospitalityDomainModule{
	{
		ID:          "HOTEL",
		ModuleName:  Hotel,
		Description: "Comprehensive hotel management module for handling reservations, room inventory, guest check-ins/outs, housekeeping, billing, and reporting.",
		IsAvailable: true,
	},
	{
		ID:          "RESTAURANT",
		ModuleName:  Restaurant,
		Description: "End-to-end restaurant operations module for managing tables, menus, orders, kitchen display systems (KDS), billing, and online ordering integration.",
		IsAvailable: true,
	},
	{
		ID:          "NIGHT_CLUB",
		ModuleName:  NightClub,
		Description: "Entertainment and nightlife operations module covering guest entry, table and bottle service, membership management, event scheduling, and POS integration.",
		IsAvailable: false,
	},
	{
		ID:          "BAR",
		ModuleName:  Bar,
		Description: "Bar and lounge management system focusing on inventory tracking, liquor control, bartender sales monitoring, and daily revenue insights.",
		IsAvailable: true,
	},
	{
		ID:          "CLOUD_KITCHEN_TAKEAWAYS",
		ModuleName:  CloudKitchen_Takeaways,
		Description: "Specialized cloud kitchen and takeaway module designed for multi-brand virtual kitchens—supports order aggregation, delivery partner APIs, kitchen workflows, and real-time analytics.",
		IsAvailable: false,
	},
}

type HospitalityDomainInfraModule struct {
	ID          string
	ModuleName  string
	Description string
	IsAvailable bool
}

var HospitalityDomainInfraModules = []HospitalityDomainInfraModule{
	{
		ID:          "8789",
		ModuleName:  "Kitchen/Service Point",
		Description: "Centralized area for food preparation, order staging, and service coordination.",
		IsAvailable: true,
	},

	{
		ID:          "8756RQ",
		ModuleName:  "Bar Counter",
		Description: "Dedicated module for beverage preparation, mixology, and direct guest service.",
		IsAvailable: false,
	},
	{
		ID:          "9910DH",
		ModuleName:  "Dining Hall",
		Description: "Manages physical floor plans, table assignments, and guest seating capacity.",
		IsAvailable: false,
	},
	{
		ID:          "4421ST",
		ModuleName:  "Inventory Store",
		Description: "Storage management for bulk supplies, dry goods, and non-perishable stock.",
		IsAvailable: false,
	},
	{
		ID:          "1022FD",
		ModuleName:  "Reception/Front Desk",
		Description: "Primary guest interface for bookings, check-ins, and concierge services.",
		IsAvailable: false,
	},
	{
		ID:          "5530HK",
		ModuleName:  "Housekeeping",
		Description: "Operational module for facility maintenance, room cleaning, and laundry services.",
		IsAvailable: false,
	},
}
