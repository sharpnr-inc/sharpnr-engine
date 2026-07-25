package general_infra_utils

// GeneralInfraModulesType enumerates the shared, business-agnostic
// infrastructure points common across venues.
type GeneralInfraModulesType string

const (
	FrontDesk      GeneralInfraModulesType = "Front Desk / Reception"
	BillingCounter GeneralInfraModulesType = "Billing Counter"
	SecurityDesk   GeneralInfraModulesType = "Security Desk"
)

type GeneralInfraModule struct {
	ID          string                  `json:"id"`
	ModuleName  GeneralInfraModulesType `json:"moduleName"`
	Description string                  `json:"description"`
	IsAvailable bool                    `json:"isAvailable"`
	Badge       string                  `json:"badge,omitempty"` // e.g. "Required", "Recommended"
	Tags        []string                `json:"tags"`
}

var GeneralInfraModules = []GeneralInfraModule{
	{
		ID:          "1022FD",
		ModuleName:  FrontDesk,
		Description: "Guest-facing desk for check-in, enquiries and walk-ins.",
		IsAvailable: true,
		Badge:       "Recommended",
		Tags:        []string{"Reception", "Walk-ins"},
	},
	{
		ID:          "2044BC",
		ModuleName:  BillingCounter,
		Description: "Cashier point for invoicing, payments and settlements.",
		IsAvailable: true,
		Badge:       "Recommended",
		Tags:        []string{"Cashier", "Payments"},
	},
	{
		ID:          "3066SD",
		ModuleName:  SecurityDesk,
		Description: "Entrance control, visitor logs and premises monitoring.",
		IsAvailable: true,
		Tags:        []string{"Entrance", "Visitors"},
	},
}

// IsValid reports whether the infra module type is a recognised value.
func (m GeneralInfraModulesType) IsValid() bool {
	switch m {
	case FrontDesk, BillingCounter, SecurityDesk:
		return true
	default:
		return false
	}
}
