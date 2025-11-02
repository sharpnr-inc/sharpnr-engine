package sharpnruirenderer

// SubSectionAccess represents access permissions for subsections within a section.
// Equivalent to: { [key: string]: boolean }
type SubSectionAccess map[string]bool

// Section represents a specific section with its access flag and subsections.
type Section struct {
	HasAccess        bool             `json:"has_access"`
	SubSectionAccess SubSectionAccess `json:"sub_section_access"`
}

// SubModuleSection represents the mapping of section names to their details.
// Equivalent to: { [key: string]: Section }
type SubModuleSection map[string]Section

// Module represents a main module (like "sharp_hospitality", "sharp_accounting", etc.).
type Module struct {
	HasAccess        bool             `json:"has_access"`
	SubModuleSection SubModuleSection `json:"sub_module_section"`
}

// MainMenuRenderTree holds all the main modules in the UI render tree.
// Equivalent to: { [key: string]: Module }
type MainMenuRenderTree map[string]Module

// UIRendererTree represents the root structure for the UI renderer tree.
type UIRendererTree struct {
	OrganizationID     string             `json:"organization_id"`
	UIRendererTreeID   string             `json:"ui_renderer_tree_id"`
	OrganizationType   string             `json:"organization_business_type"`
	MainMenuRenderTree MainMenuRenderTree `json:"main_menu_render_tree"`
}
