package healthcare_utilities

// HealthcareModulesType enumerates the patient-facing service lines a
// healthcare provider can operate.
//
// The infrastructure modules that support these service lines live in
// sharpnr-engine/property_infra/healthcare_infra/utils.
type HealthcareModulesType string

const (
	Hospital     HealthcareModulesType = "HOSPITAL"
	Clinic       HealthcareModulesType = "CLINIC"
	Pharmacy     HealthcareModulesType = "PHARMACY"
	Diagnostics  HealthcareModulesType = "DIAGNOSTICS"
	Telemedicine HealthcareModulesType = "TELEMEDICINE"
	HomeCare     HealthcareModulesType = "HOME_CARE"
)

// Business module
type HealthcareModule struct {
	ID          string                `json:"id"`
	ModuleName  HealthcareModulesType `json:"moduleName"`
	Description string                `json:"description"`
	IsAvailable bool                  `json:"isAvailable"`
	Tags        []string              `json:"tags"`
}

var HealthcareModules = []HealthcareModule{
	{
		ID:          "HOSPITAL",
		ModuleName:  Hospital,
		Description: "Comprehensive hospital management module covering patient admissions, ward and bed management, OPD/IPD workflows, billing, and clinical records.",
		IsAvailable: false,
		Tags:        []string{"patients", "admissions", "ipd", "opd", "billing"},
	},
	{
		ID:          "CLINIC",
		ModuleName:  Clinic,
		Description: "Outpatient clinic module for appointment scheduling, consultations, prescriptions, and patient follow-up management.",
		IsAvailable: false,
		Tags:        []string{"appointments", "consultations", "prescriptions"},
	},
	{
		ID:          "PHARMACY",
		ModuleName:  Pharmacy,
		Description: "Pharmacy management module handling drug inventory, dispensing, prescription validation, and stock expiry tracking.",
		IsAvailable: false,
		Tags:        []string{"drugs", "inventory", "dispensing", "stock"},
	},
	{
		ID:          "DIAGNOSTICS",
		ModuleName:  Diagnostics,
		Description: "Diagnostics and laboratory module for managing test orders, sample tracking, result reporting, and lab equipment integration.",
		IsAvailable: false,
		Tags:        []string{"lab", "tests", "samples", "reports"},
	},
	{
		ID:          "TELEMEDICINE",
		ModuleName:  Telemedicine,
		Description: "Remote consultation module supporting video visits, e-prescriptions, secure messaging, and virtual follow-ups.",
		IsAvailable: false,
		Tags:        []string{"video", "remote", "e-prescription", "virtual"},
	},
	{
		ID:          "HOME_CARE",
		ModuleName:  HomeCare,
		Description: "Home healthcare module for scheduling visiting staff, managing care plans, and tracking at-home patient vitals.",
		IsAvailable: false,
		Tags:        []string{"home-visit", "care-plan", "nursing", "vitals"},
	},
}

// ---------------------------------------------------------------------------
// Utilities
// ---------------------------------------------------------------------------

// IsValid reports whether the business module type is a recognised value.
func (m HealthcareModulesType) IsValid() bool {
	switch m {
	case Hospital, Clinic, Pharmacy, Diagnostics, Telemedicine, HomeCare:
		return true
	default:
		return false
	}
}

// GetModuleByID returns the business module with the given ID.
// The second return value is false if no module matches.
func GetModuleByID(id string) (HealthcareModule, bool) {
	for _, m := range HealthcareModules {
		if m.ID == id {
			return m, true
		}
	}
	return HealthcareModule{}, false
}

// GetModuleByType returns the business module with the given module type.
func GetModuleByType(t HealthcareModulesType) (HealthcareModule, bool) {
	for _, m := range HealthcareModules {
		if m.ModuleName == t {
			return m, true
		}
	}
	return HealthcareModule{}, false
}

// GetAvailableModules returns only the business modules flagged available.
func GetAvailableModules() []HealthcareModule {
	out := make([]HealthcareModule, 0, len(HealthcareModules))
	for _, m := range HealthcareModules {
		if m.IsAvailable {
			out = append(out, m)
		}
	}
	return out
}

// GetModulesByTag returns business modules carrying the given tag.
func GetModulesByTag(tag string) []HealthcareModule {
	var out []HealthcareModule
	for _, m := range HealthcareModules {
		for _, t := range m.Tags {
			if t == tag {
				out = append(out, m)
				break
			}
		}
	}
	return out
}
