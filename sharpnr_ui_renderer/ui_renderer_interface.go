package sharpnruirenderer

import (
	"context"
	"sharpnr-engine/core/models"
)

type UIRendererRepository interface {
	CreateOrgUIRenderer(ctx context.Context, renderer *models.OrgUIRendererTree) error
	GetOrgUIRenderer()
	UpdateOrgUIRenderer()
}
