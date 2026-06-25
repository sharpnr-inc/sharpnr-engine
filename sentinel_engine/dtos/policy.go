package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// AddPolicyRequestDto is the create-policy payload. The frontend builds and
// validates the Cedar text (via @cedar-policy/cedar-wasm) and sends it here as a
// string; the service revalidates before persisting.
type AddPolicyRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Policy      string `json:"policy"`  // raw Cedar policy text from the frontend
	IsEnabled   bool   `json:"enabled"` // whether the policy is active in the engine
}

// PolicyView is the read shape for a stored policy. It exposes the raw Cedar text
// (the FE renders/edits it with cedar-wasm) alongside its admin metadata.
type PolicyView struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Policy      string    `json:"policy"`
	Enabled     bool      `json:"enabled"`
}

// ToPolicyView projects a policy model to its read shape.
func ToPolicyView(p *sentinelengine_models.SentinelPolicy) PolicyView {
	return PolicyView{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Policy:      p.Policy,
		Enabled:     p.Enabled,
	}
}

// ToPolicyViews projects a slice of policy models to read shapes.
func ToPolicyViews(ps []*sentinelengine_models.SentinelPolicy) []PolicyView {
	out := make([]PolicyView, 0, len(ps))
	for _, p := range ps {
		out = append(out, ToPolicyView(p))
	}
	return out
}
