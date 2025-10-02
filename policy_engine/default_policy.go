package policyengine

// __METADATA__ section
type Metadata struct {
	OrgID     string  `json:"ORG_ID"`
	SharpnrID string  `json:"SHARPNR_ID"`
	PolicyID  string  `json:"POLICY_ID"`
	Version   float64 `json:"VERSION"`
	OrgType   string  `json:"ORG_TYPE"`
	// Modules   []interface{} `json:"MODULES"` // mixed: string + object
}

type OrgDefaultPolicy struct {
	Metadata Metadata `json:"__METADATA__"`
}
