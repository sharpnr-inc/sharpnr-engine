package templates

import sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"

func HospitalityTemplate() sharpnruirenderer.RenderTree {
	return sharpnruirenderer.RenderTree{
		"sharp_hospitality": {
			HasAccess: true,
			SubModuleSection: map[string]sharpnruirenderer.RenderNode{
				"bar": {
					HasAccess: true,
					SubSectionAccess: map[string]bool{
						"menu": true,
					},
				},

				"restaurant": {
					HasAccess: true,
					SubSectionAccess: map[string]bool{
						"order": true,
					},
				},

				"hotel": {
					HasAccess: true,
					SubSectionAccess: map[string]bool{
						"booking":         true,
						"room_management": true,
					},
				},

				"kitchen": {
					HasAccess: true,
				},
				"inventory": {
					HasAccess: true,
				},
				"settings": {
					HasAccess: true,
				},
			},
		},
	}
}
