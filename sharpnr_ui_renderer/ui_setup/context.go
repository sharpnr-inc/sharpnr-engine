package uisetup

import (
	"sharpnr-engine/core/models"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"
)

type SetupContext struct {
	Org          *models.Organization
	Modules      []string
	UIRenderRepo sharpnruirenderer.UIRendererRepository
}
