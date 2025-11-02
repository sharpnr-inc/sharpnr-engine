package sharpnruirenderer

import (
	"context"
	"sharpnr-engine/core/constants"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func (o *orgUIRendererRepository) CreateOrgUIRenderer(ctx context.Context, orgId uuid.UUID, orgConfig []string, orgType constants.OrgType) (datatypes.JSON, error) {

	panic("unimplemented")
}
