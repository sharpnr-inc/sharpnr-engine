package hospitality_utilities

type HospitalityDomainModule struct {
	ID          string
	ModuleName  string
	Description string
	IsAvailable bool
}

var HospitalityModules = []HospitalityDomainModule{
	{
		ID:          "8908",
		ModuleName:  "HOTEL",
		Description: "Comprehensive hotel management module for handling reservations, room inventory, guest check-ins/outs, housekeeping, billing, and reporting.",
		IsAvailable: false,
	},
	{
		ID:          "3453",
		ModuleName:  "RESTAURANT",
		Description: "End-to-end restaurant operations module for managing tables, menus, orders, kitchen display systems (KDS), billing, and online ordering integration.",
		IsAvailable: true,
	},
	{
		ID:          "3453",
		ModuleName:  "PUB/CLUB",
		Description: "Entertainment and nightlife operations module covering guest entry, table and bottle service, membership management, event scheduling, and POS integration.",
		IsAvailable: false,
	},
	{
		ID:          "8394758",
		ModuleName:  "BAR",
		Description: "Bar and lounge management system focusing on inventory tracking, liquor control, bartender sales monitoring, and daily revenue insights.",
		IsAvailable: true,
	},
	{
		ID:          "heiu4533",
		ModuleName:  "CLOUD KITCHEN/TAKEAWAYS",
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
		IsAvailable: true,
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
