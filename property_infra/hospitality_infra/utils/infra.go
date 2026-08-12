package hospitality_infra_utils

// HospitalityInfraModulesTypes enumerates the operational / physical
// infrastructure areas that support the hospitality business modules.
type HospitalityInfraModulesTypes string

const (
	Kitchen        HospitalityInfraModulesTypes = "KITCHEN"
	BarCounter     HospitalityInfraModulesTypes = "BAR_COUNTER"
	DiningHall     HospitalityInfraModulesTypes = "DINING_HALL"
	InventoryStore HospitalityInfraModulesTypes = "INVENTORY_STORE"
	FrontDesk      HospitalityInfraModulesTypes = "RECEPTION_FRONT_DESK"
	Housekeeping   HospitalityInfraModulesTypes = "HOUSEKEEPING"
)

// Infra module
type HospitalityInfraModule struct {
	ID          string                       `json:"id"`
	ModuleName  HospitalityInfraModulesTypes `json:"moduleName"`
	Description string                       `json:"description"`
	IsAvailable bool                         `json:"isAvailable"`
	Tags        []string                     `json:"tags"`
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

// IsValid reports whether the infra module type is a recognised value.
func (m HospitalityInfraModulesTypes) IsValid() bool {
	switch m {
	case Kitchen, BarCounter, DiningHall, InventoryStore, FrontDesk, Housekeeping:
		return true
	default:
		return false
	}
}

// GetInfraModuleByID returns the infra module with the given ID.
// The second return value is false if no module matches.
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
