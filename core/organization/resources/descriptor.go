// Package organizationResources declares the "Organization" sentinel resource:
// the org itself, the resource that org-level administration actions act on
// (managing authorization, and later org settings). It registers only a
// Descriptor (actions + attributes), no instance Provider — there is exactly one
// organization per tenant (the current org), so the policy builder scopes admin
// policies to it implicitly; there is no instance list to fetch.
package organizationResources

import (
	organizationSentinalActions "sharpnr-engine/core/organization/actions"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

// ResourceType is the Cedar entity type for the organization. It uses the
// engine's exported constant so the builder and the evaluator never drift.
const ResourceType = sentinelengine.EntityOrganization // "Organization"

var Descriptor = sentinel_resources.Descriptor{
	Type:        ResourceType,
	Label:       "Organization",
	Description: "The organization itself — the resource for org-level administration (managing authorization).",
	Actions:     organizationSentinalActions.Actions,
	Attributes:  nil, // orgId is stamped on every resource automatically by the engine
}
