package sharpnruirenderer

type RenderNode struct {
	HasAccess        bool                  `json:"has_access"`
	SubModuleSection map[string]RenderNode `json:"sub_module_section,omitempty"`
	SubSectionAccess map[string]bool       `json:"sub_section_access,omitempty"`
}

type RenderTree map[string]RenderNode
