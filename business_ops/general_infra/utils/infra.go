package general_infra_utils

type GeneralInfraModule struct {
	ID          string   `json:"id"`
	ModuleName  string   `json:"moduleName"`
	Description string   `json:"description"`
	IsAvailable bool     `json:"isAvailable"`
	Badge       string   `json:"badge,omitempty"` // e.g. "Required", "Recommended"
	Tags        []string `json:"tags"`
}

var GeneralInfraModules = []GeneralInfraModule{
	{
		ID:          "1022FD",
		ModuleName:  "Front Desk / Reception",
		Description: "Guest-facing desk for check-in, enquiries and walk-ins.",
		IsAvailable: true,
		Badge:       "Required",
		Tags:        []string{"Reception", "Walk-ins"},
	},
	{
		ID:          "2044BC",
		ModuleName:  "Billing Counter",
		Description: "Cashier point for invoicing, payments and settlements.",
		IsAvailable: true,
		Badge:       "Recommended",
		Tags:        []string{"Cashier", "Payments"},
	},
	{
		ID:          "3066SD",
		ModuleName:  "Security Desk",
		Description: "Entrance control, visitor logs and premises monitoring.",
		IsAvailable: true,
		Tags:        []string{"Entrance", "Visitors"},
	},
}
