package sentinelservice

import (
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
)

// Catalog returns the full permission catalog — the source the frontend renders
// when an admin builds a policy. It is the aggregated registry every feature
// module contributes its actions to at startup (see sentinel_actions.Register).
func (s *sentinelService) Catalog() []sentinel_actions.Permission {
	return sentinel_actions.Catalog()
}
