package sentinelDefaultpolicies

// SentinalDefaultPolices is the Cedar text of a starter policy seeded for a
// well-known role slug. These are the "out of the box" grants an org gets for
// the executive roles created at provisioning, before an admin customises them
// through managed policies.
//
// Every policy follows the same shape the generator emits (see
// service.buildCedarPolicy): principal scoped to a Role slug, an action set
// (explicit list or the wildcard), and the tenant guard
// `resource.orgId == principal.orgId` so a grant can never reach another org.
// Action lists are alphabetised to match the generator's byte-stable output.
//
// Action ids must exist in the actions catalog (sentinel_actions). Keep these in
// sync with core_action.go.
type SentinalDefaultPolices string

const (
	// OwnerPolcy: full access within the org. The action position is left
	// unconstrained (the wildcard), so it also covers actions added later.
	OwnerPolcy SentinalDefaultPolices = `permit (
  principal in Role::"owner",
  action,
  resource
) when { resource.orgId == principal.orgId };`
)
