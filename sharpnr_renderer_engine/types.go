package sharpnr_renderer_engine

// RenderNode represents a UI node
type RenderNode struct {
	HasAccess        bool                  `json:"has_access"`
	SubModuleSection map[string]RenderNode `json:"sub_module_section,omitempty"`
	SubSectionAccess map[string]bool       `json:"sub_section_access,omitempty"`
}

// RenderTree is the full UI tree sent to frontend
type RenderTree map[string]RenderNode
