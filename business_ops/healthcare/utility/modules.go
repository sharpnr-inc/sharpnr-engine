package healthcare_utilities

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

// Infra module
type HealthcareInfraModule struct {
	ID          string   `json:"id"`
	ModuleName  string   `json:"moduleName"`
	Description string   `json:"description"`
	IsAvailable bool     `json:"isAvailable"`
	Tags        []string `json:"tags"`
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
		ModuleName:  "Ward/Inpatient Unit",
		Description: "Manages bed allocation, ward assignments, and inpatient capacity planning.",
		IsAvailable: false,
		Tags:        []string{"wards", "beds", "ipd"},
	},
	{
		ID:          "7745OT",
		ModuleName:  "Operation Theatre",
		Description: "Surgical suite module for OT scheduling, equipment readiness, and procedure tracking.",
		IsAvailable: false,
		Tags:        []string{"surgery", "ot", "scheduling"},
	},
	{
		ID:          "7720LB",
		ModuleName:  "Laboratory",
		Description: "Physical lab space for sample processing, testing, and result generation.",
		IsAvailable: false,
		Tags:        []string{"lab", "samples", "testing"},
	},
	{
		ID:          "7733PH",
		ModuleName:  "Pharmacy Counter",
		Description: "Dispensing point for medication handover and prescription fulfillment.",
		IsAvailable: false,
		Tags:        []string{"pharmacy", "dispensing", "medication"},
	},
	{
		ID:          "7710RC",
		ModuleName:  "Reception/Registration",
		Description: "Primary patient interface for registration, appointment check-in, and queue management.",
		IsAvailable: false,
		Tags:        []string{"registration", "front-desk", "queue"},
	},
	{
		ID:          "7758RD",
		ModuleName:  "Radiology/Imaging",
		Description: "Imaging unit for X-ray, CT, MRI, and ultrasound services with report integration.",
		IsAvailable: false,
		Tags:        []string{"imaging", "radiology", "scans"},
	},
}
