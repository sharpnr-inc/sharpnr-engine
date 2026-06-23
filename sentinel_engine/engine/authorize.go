package sentinelengine

import (
	"context"
	"fmt"

	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
)

// IsAuthorized answers a single request: can this user perform this action on
// this resource, within this org? It loads (or reuses) the org's compiled
// policies, builds the principal+resource entity graph from current membership,
// and evaluates. On any infrastructure error it returns (deny, err) — fail closed.
func (s *Sentinel) IsAuthorized(ctx context.Context, in AuthorizeInput) (Result, error) {
	ps, err := s.policySetFor(ctx, in.OrgID)
	if err != nil {
		return Result{Allowed: false}, fmt.Errorf("load policies: %w", err)
	}

	entities, err := s.buildEntities(ctx, in)
	if err != nil {
		return Result{Allowed: false}, fmt.Errorf("build entities: %w", err)
	}

	req := cedar.Request{
		Principal: cedar.NewEntityUID(entityUser, types.String(in.UserID.String())),
		Action:    sentinel_actions.Action(in.Action),
		Resource:  cedar.NewEntityUID(types.EntityType(in.Resource.Type), types.String(in.Resource.ID)),
		Context:   cedar.NewRecord(toRecordMap(in.Context)),
	}

	decision, diag := cedar.Authorize(ps, entities, req)

	res := Result{Allowed: decision == cedar.Allow}
	for _, r := range diag.Reasons {
		res.Reasons = append(res.Reasons, string(r.PolicyID))
	}
	for _, de := range diag.Errors {
		res.Errors = append(res.Errors, fmt.Sprintf("%s: %s", de.PolicyID, de.Message))
	}
	return res, nil
}

// toRecordMap converts a plain string-keyed value map into a Cedar RecordMap,
// used to build the request Context.
func toRecordMap(m map[string]types.Value) cedar.RecordMap {
	rm := cedar.RecordMap{}
	for k, v := range m {
		rm[types.String(k)] = v
	}
	return rm
}
