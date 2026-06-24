package defaultpolicies

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

	// CEOPolicy: top executive. Full operational access across the org, same
	// wildcard reach as the owner — the distinction is organisational (account
	// owner vs. chief executive), not a difference in what Cedar permits.
	CEOPolicy SentinalDefaultPolices = `permit (
  principal in Role::"ceo",
  action,
  resource
) when { resource.orgId == principal.orgId };`

	// CTOPolicy: technical leadership. Reads the repo and the org's properties,
	// and has full control over documents. Not granted property mutation
	// (creating properties / adding business modules) — that's an operations
	// concern (see COOPolicy).
	CTOPolicy SentinalDefaultPolices = `permit (
  principal in Role::"cto",
  action in [Action::"deleteDocument", Action::"editDocument", Action::"viewDocument", Action::"viewProperty", Action::"viewRepo"],
  resource
) when { resource.orgId == principal.orgId };`

	// COOPolicy: operations. Full control over properties (view, create, add
	// business modules) and day-to-day document work (view, edit). Not granted
	// document deletion or repo access.
	COOPolicy SentinalDefaultPolices = `permit (
  principal in Role::"coo",
  action in [Action::"addBusinessModulesInProperty", Action::"createProperty", Action::"editDocument", Action::"viewDocument", Action::"viewProperty"],
  resource
) when { resource.orgId == principal.orgId };`
)
