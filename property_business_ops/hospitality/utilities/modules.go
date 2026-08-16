package hospitality_utilities

// HospitalityModulesType enumerates the customer-facing business lines a
// hospitality venue can operate.
//
// The infrastructure modules that support these business lines live in
// sharpnr-engine/property_infra/hospitality_infra/utils.
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
