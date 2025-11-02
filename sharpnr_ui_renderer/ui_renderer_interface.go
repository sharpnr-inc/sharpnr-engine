package sharpnruirenderer

import (
	"context"

	"gorm.io/datatypes"
)

type UIRendererRepository interface {
	CreateOrgUIRenderer(ctx context.Context, orgId string) (datatypes.JSON, error)
	GetOrgUIRenderer()
	UpdateOrgUIRenderer()
}
