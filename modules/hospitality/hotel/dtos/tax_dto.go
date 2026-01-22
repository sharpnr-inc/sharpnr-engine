package hoteldtos

import (
	"time"
)

type CreateHotelTaxRuleRequest struct {
	RuleName      string `json:"rule_name"`
	TaxType       string `json:"tax_type"`
	ConditionType string `json:"condition_type"`

	MinValue *float64 `json:"min_value,omitempty"`
	MaxValue *float64 `json:"max_value,omitempty"`

	TaxRate float64 `json:"tax_rate"`

	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`

	Components []CreateHotelTaxComponent `json:"components"`
}

type CreateHotelTaxComponent struct {
	ComponentType string  `json:"component_type"`
	Rate          float64 `json:"rate"`
}
