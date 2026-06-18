package hospitality_utilities

// HospitalityModulesType enumerates the customer-facing business lines a
// hospitality venue can operate.
type HospitalityModulesType string

const (
	Restaurant             HospitalityModulesType = "RESTAURANT"
	Hotel                  HospitalityModulesType = "HOTEL"
	Bar                    HospitalityModulesType = "BAR"
	NightClub              HospitalityModulesType = "NIGHT_CLUB"
	Resort                 HospitalityModulesType = "RESORT"
	CloudKitchen_Takeaways HospitalityModulesType = "CLOUD_KITCHEN/TAKEAWAYS"
)

// HospitalityInfraModulesTypes enumerates the operational / physical
// infrastructure areas that support the business modules above.
type HospitalityInfraModulesTypes string

const (
	Kitchen        HospitalityInfraModulesTypes = "KITCHEN"
	BarCounter     HospitalityInfraModulesTypes = "BAR_COUNTER"
	DiningHall     HospitalityInfraModulesTypes = "DINING_HALL"
	InventoryStore HospitalityInfraModulesTypes = "INVENTORY_STORE"
	FrontDesk      HospitalityInfraModulesTypes = "RECEPTION_FRONT_DESK"
	Housekeeping   HospitalityInfraModulesTypes = "HOUSEKEEPING"
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
	ID          string                       `json:"id"`
	ModuleName  HospitalityInfraModulesTypes `json:"moduleName"`
	Description string                       `json:"description"`
	IsAvailable bool                         `json:"isAvailable"`
	Tags        []string                     `json:"tags"`
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
		ID:          "RESORT",
		ModuleName:  Resort,
		Description: "Full-service resort management module combining accommodation, dining, recreation, spa, and activity scheduling under a single consolidated guest folio with unified billing and reporting.",
		IsAvailable: false,
		Tags:        []string{"accommodation", "leisure", "recreation", "spa", "billing"},
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
		ModuleName:  Kitchen,
		Description: "Centralized area for food preparation, order staging, and service coordination.",
		IsAvailable: true,
		Tags:        []string{"kitchen", "food-prep", "service"},
	},
	{
		ID:          "8756RQ",
		ModuleName:  BarCounter,
		Description: "Dedicated module for beverage preparation, mixology, and direct guest service.",
		IsAvailable: false,
		Tags:        []string{"bar", "beverages", "service"},
	},
	{
		ID:          "9910DH",
		ModuleName:  DiningHall,
		Description: "Manages physical floor plans, table assignments, and guest seating capacity.",
		IsAvailable: false,
		Tags:        []string{"seating", "floor-plan", "tables"},
	},
	{
		ID:          "4421ST",
		ModuleName:  InventoryStore,
		Description: "Storage management for bulk supplies, dry goods, and non-perishable stock.",
		IsAvailable: false,
		Tags:        []string{"inventory", "storage", "supplies"},
	},
	{
		ID:          "1022FD",
		ModuleName:  FrontDesk,
		Description: "Primary guest interface for bookings, check-ins, and concierge services.",
		IsAvailable: false,
		Tags:        []string{"front-desk", "bookings", "concierge"},
	},
	{
		ID:          "5530HK",
		ModuleName:  Housekeeping,
		Description: "Operational module for facility maintenance, room cleaning, and laundry services.",
		IsAvailable: false,
		Tags:        []string{"housekeeping", "cleaning", "maintenance"},
	},
}

// ---------------------------------------------------------------------------
// Utilities
// ---------------------------------------------------------------------------

// IsValid reports whether the business module type is a recognised value.
func (m HospitalityModulesType) IsValid() bool {
	switch m {
	case Restaurant, Hotel, Bar, NightClub, Resort, CloudKitchen_Takeaways:
		return true
	default:
		return false
	}
}

// IsValid reports whether the infra module type is a recognised value.
func (m HospitalityInfraModulesTypes) IsValid() bool {
	switch m {
	case Kitchen, BarCounter, DiningHall, InventoryStore, FrontDesk, Housekeeping:
		return true
	default:
		return false
	}
}

// GetModuleByID returns the business module with the given ID.
// The second return value is false if no module matches.
func GetModuleByID(id string) (HospitalityModule, bool) {
	for _, m := range HospitalityModules {
		if m.ID == id {
			return m, true
		}
	}
	return HospitalityModule{}, false
}

// GetModuleByType returns the business module with the given module type.
func GetModuleByType(t HospitalityModulesType) (HospitalityModule, bool) {
	for _, m := range HospitalityModules {
		if m.ModuleName == t {
			return m, true
		}
	}
	return HospitalityModule{}, false
}

// GetInfraModuleByID returns the infra module with the given ID.
func GetInfraModuleByID(id string) (HospitalityInfraModule, bool) {
	for _, m := range HospitalityInfraModules {
		if m.ID == id {
			return m, true
		}
	}
	return HospitalityInfraModule{}, false
}

// GetInfraModuleByType returns the infra module with the given module type.
func GetInfraModuleByType(t HospitalityInfraModulesTypes) (HospitalityInfraModule, bool) {
	for _, m := range HospitalityInfraModules {
		if m.ModuleName == t {
			return m, true
		}
	}
	return HospitalityInfraModule{}, false
}

// GetAvailableModules returns only the business modules flagged available.
func GetAvailableModules() []HospitalityModule {
	out := make([]HospitalityModule, 0, len(HospitalityModules))
	for _, m := range HospitalityModules {
		if m.IsAvailable {
			out = append(out, m)
		}
	}
	return out
}

// GetAvailableInfraModules returns only the infra modules flagged available.
func GetAvailableInfraModules() []HospitalityInfraModule {
	out := make([]HospitalityInfraModule, 0, len(HospitalityInfraModules))
	for _, m := range HospitalityInfraModules {
		if m.IsAvailable {
			out = append(out, m)
		}
	}
	return out
}

// GetModulesByTag returns business modules carrying the given tag.
func GetModulesByTag(tag string) []HospitalityModule {
	var out []HospitalityModule
	for _, m := range HospitalityModules {
		for _, t := range m.Tags {
			if t == tag {
				out = append(out, m)
				break
			}
		}
	}
	return out
}

// GetInfraModulesByTag returns infra modules carrying the given tag.
func GetInfraModulesByTag(tag string) []HospitalityInfraModule {
	var out []HospitalityInfraModule
	for _, m := range HospitalityInfraModules {
		for _, t := range m.Tags {
			if t == tag {
				out = append(out, m)
				break
			}
		}
	}
	return out
}
