package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel/models"
	"time"

	"github.com/google/uuid"
)

type CreateGroupsRequestDto struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ParentGroupID *uuid.UUID `json:"parentGroupId"` // nullable: top-level groups have no parent
}

// GroupView is the read shape for a group in the group tree.
type GroupView struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ParentGroupID *uuid.UUID `json:"parentGroupId"`
}

// GroupDetailView is what the SENTINEL service can answer on its own: the group,
// the policies registered against it, and the ids of its members.
//
// Members are ids, not names. Employee names live in the employee module, and
// sentinel deliberately has no dependency on it — user_groups.user_id is an
// employee id with no FK. The orchestrator composes the two and produces
// GroupDetailsResponseDto below.
type GroupDetailView struct {
	Group     GroupView
	CreatedAt time.Time
	UpdatedAt time.Time
	Policies  []PolicyView
	MemberIDs []uuid.UUID
}

/* ---------------------------------- <***> --------------------------------- */

// GroupDetailsResponseDto is the wire shape: GroupDetailView with member ids
// resolved to names by the orchestrator.
type GroupDetailsResponseDto struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        []GroupUserResponseDto    // list of all user in the Group
	Polices     []GroupPolicesResponseDto // List of polices in the Group
}

type GroupUserResponseDto struct {
	ID   uuid.UUID
	Name string
}

type GroupPolicesResponseDto struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
}

// ToGroupView projects a group model to its read shape.
func ToGroupView(g *sentinelengine_models.Group) GroupView {
	return GroupView{
		ID:            g.ID,
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
