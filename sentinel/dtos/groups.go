package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

type CreateGroupsRequestDto struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Slug          string     `json:"slug"`
	ParentGroupID *uuid.UUID `json:"parentGroupId"` // nullable: top-level groups have no parent
}

// GroupView is the read shape for a group in the group tree.
type GroupView struct {
	ID            uuid.UUID  `json:"id"`
	Slug          string     `json:"slug"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ParentGroupID *uuid.UUID `json:"parentGroupId"`
}

// ToGroupView projects a group model to its read shape.
func ToGroupView(g *sentinelengine_models.Group) GroupView {
	return GroupView{
		ID:            g.ID,
		Slug:          g.Slug,
		Name:          g.Name,
		Description:   g.Description,
		ParentGroupID: g.ParentGroupID,
	}
}

// ToGroupViews projects a slice of group models to read shapes.
func ToGroupViews(gs []*sentinelengine_models.Group) []GroupView {
	out := make([]GroupView, 0, len(gs))
	for _, g := range gs {
		out = append(out, ToGroupView(g))
	}
	return out
}
