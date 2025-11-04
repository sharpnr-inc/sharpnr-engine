package sharpnruirenderer

import "github.com/google/uuid"

// ---- Root Object ----
// This is Render Tree UI Top Level
type UIRendererTree struct {
	OrganizationID           uuid.UUID          `json:"organization_id"`
	UIRendererTreeID         uuid.UUID          `json:"ui_renderer_tree_id"`
	OrganizationBusinessType string             `json:"organization_business_type"`
	MainMenuRenderTree       MainMenuRenderTree `json:"main_menu_render_tree"`
}

// ---- Main Menu Render Tree ----
type MainMenuRenderTree struct {
	SharpHospitality SharpHospitalityTree `json:"sharp_hospitality"`
	SharpAccounting  SharpAccountingTree  `json:"sharp_accounting"`
	SharpHR          SharpHRTree          `json:"sharp_hr"`
	SharpMarketing   SharpMarketingTree   `json:"sharp_marketing"`
	Employee         EmployeeTree         `json:"employee"`
}
