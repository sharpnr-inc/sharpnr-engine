package healthcare_utilities

// HealthcareModulesType enumerates the patient-facing service lines a
// healthcare provider can operate.
type HealthcareModulesType string

const (
	Hospital     HealthcareModulesType = "HOSPITAL"
	Clinic       HealthcareModulesType = "CLINIC"
	Pharmacy     HealthcareModulesType = "PHARMACY"
	Diagnostics  HealthcareModulesType = "DIAGNOSTICS"
	Telemedicine HealthcareModulesType = "TELEMEDICINE"
	HomeCare     HealthcareModulesType = "HOME_CARE"
)

// HealthcareInfraModulesType enumerates the operational / physical
// infrastructure areas that support the business modules above.
type HealthcareInfraModulesType string

const (
	Ward             HealthcareInfraModulesType = "WARD_INPATIENT_UNIT"
	OperationTheatre HealthcareInfraModulesType = "OPERATION_THEATRE"
	Laboratory       HealthcareInfraModulesType = "LABORATORY"
	PharmacyCounter  HealthcareInfraModulesType = "PHARMACY_COUNTER"
	Reception        HealthcareInfraModulesType = "RECEPTION_REGISTRATION"
	Radiology        HealthcareInfraModulesType = "RADIOLOGY_IMAGING"
)

// Business module
type HealthcareModule struct {
	ID          string                `json:"id"`
	ModuleName  HealthcareModulesType `json:"moduleName"`
	Description string                `json:"description"`
	IsAvailable bool                  `json:"isAvailable"`
	Tags        []string              `json:"tags"`
}

// Infra module
type HealthcareInfraModule struct {
	ID          string                     `json:"id"`
	ModuleName  HealthcareInfraModulesType `json:"moduleName"`
	Description string                     `json:"description"`
	IsAvailable bool                       `json:"isAvailable"`
	Tags        []string                   `json:"tags"`
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

var HealthcareInfraModules = []HealthcareInfraModule{
	{
		ID:          "7701WD",
		ModuleName:  Ward,
		Description: "Manages bed allocation, ward assignments, and inpatient capacity planning.",
		IsAvailable: false,
		Tags:        []string{"wards", "beds", "ipd"},
	},
	{
		ID:          "7745OT",
		ModuleName:  OperationTheatre,
		Description: "Surgical suite module for OT scheduling, equipment readiness, and procedure tracking.",
		IsAvailable: false,
		Tags:        []string{"surgery", "ot", "scheduling"},
	},
	{
		ID:          "7720LB",
		ModuleName:  Laboratory,
		Description: "Physical lab space for sample processing, testing, and result generation.",
		IsAvailable: false,
		Tags:        []string{"lab", "samples", "testing"},
	},
	{
		ID:          "7733PH",
		ModuleName:  PharmacyCounter,
		Description: "Dispensing point for medication handover and prescription fulfillment.",
		IsAvailable: false,
		Tags:        []string{"pharmacy", "dispensing", "medication"},
	},
	{
		ID:          "7710RC",
		ModuleName:  Reception,
		Description: "Primary patient interface for registration, appointment check-in, and queue management.",
		IsAvailable: false,
		Tags:        []string{"registration", "front-desk", "queue"},
	},
	{
		ID:          "7758RD",
		ModuleName:  Radiology,
		Description: "Imaging unit for X-ray, CT, MRI, and ultrasound services with report integration.",
		IsAvailable: false,
		Tags:        []string{"imaging", "radiology", "scans"},
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

// IsValid reports whether the infra module type is a recognised value.
func (m HealthcareInfraModulesType) IsValid() bool {
	switch m {
	case Ward, OperationTheatre, Laboratory, PharmacyCounter, Reception, Radiology:
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

// GetInfraModuleByID returns the infra module with the given ID.
func GetInfraModuleByID(id string) (HealthcareInfraModule, bool) {
	for _, m := range HealthcareInfraModules {
		if m.ID == id {
			return m, true
		}
	}
	return HealthcareInfraModule{}, false
}

// GetInfraModuleByType returns the infra module with the given module type.
func GetInfraModuleByType(t HealthcareInfraModulesType) (HealthcareInfraModule, bool) {
	for _, m := range HealthcareInfraModules {
		if m.ModuleName == t {
			return m, true
		}
	}
	return HealthcareInfraModule{}, false
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

// GetAvailableInfraModules returns only the infra modules flagged available.
func GetAvailableInfraModules() []HealthcareInfraModule {
	out := make([]HealthcareInfraModule, 0, len(HealthcareInfraModules))
	for _, m := range HealthcareInfraModules {
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

// GetInfraModulesByTag returns infra modules carrying the given tag.
func GetInfraModulesByTag(tag string) []HealthcareInfraModule {
	var out []HealthcareInfraModule
	for _, m := range HealthcareInfraModules {
		for _, t := range m.Tags {
			if t == tag {
				out = append(out, m)
				break
			}
		}
	}
	return out
}
