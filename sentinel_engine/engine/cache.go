package sentinelengine

import (
	"context"
	"fmt"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// Policy-set cache. Policies change rarely, so each org's compiled
// *cedar.PolicySet is cached and reused across authorization calls until
// explicitly invalidated.

// InvalidateOrg drops an org's compiled PolicySet so the next authorization
// recompiles from the database. Call this after any create/update/delete of an
// org's policies. Role/group *membership* changes do NOT need this — entities
// are rebuilt on every request.
func (s *Sentinel) InvalidateOrg(orgID uuid.UUID) {
	s.mu.Lock()
	delete(s.cache, orgID)
	s.mu.Unlock()
}

// policySetFor returns the cached compiled set for an org, compiling on a miss.
func (s *Sentinel) policySetFor(ctx context.Context, orgID uuid.UUID) (*cedar.PolicySet, error) {
	s.mu.RLock()
	ps, ok := s.cache[orgID]
	s.mu.RUnlock()
	if ok {
		return ps, nil
	}

	ps, err := s.compile(ctx, orgID)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.cache[orgID] = ps
	s.mu.Unlock()
	return ps, nil
}

// compile turns an org's stored policy rows into a *cedar.PolicySet. The row's
// ID is used as the Cedar PolicyID, so diagnostics point back at your table.
func (s *Sentinel) compile(ctx context.Context, orgID uuid.UUID) (*cedar.PolicySet, error) {
	rows, err := s.repo.ListEnabledPolicies(ctx, orgID)
	if err != nil {
		return nil, err
	}

	ps := cedar.NewPolicySet()
	for _, row := range rows {
		var p cedar.Policy
		if err := p.UnmarshalCedar([]byte(row.Policy)); err != nil {
			// Fail closed: a malformed stored policy is an integrity error.
			// Policies are validated on write, so this should not happen in
			// practice. Returning here denies the whole org rather than
			// silently dropping one policy and granting more than intended.
			return nil, fmt.Errorf("policy %s failed to parse: %w", row.ID, err)
		}
		ps.Add(cedar.PolicyID(row.ID.String()), &p)
	}
	return ps, nil
}
