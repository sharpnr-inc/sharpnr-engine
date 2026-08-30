package sentinelOwnResources

import (
	"context"
	"strings"

	agg "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// The instance pickers.
//
// All three filter in memory over the service's existing org-scoped list calls.
// That is honest for the sizes involved — an org has tens of roles and groups,
// not thousands — but it is the wrong shape once a list grows: the filter should
// move into the service and then into SQL. The signature is already right for
// that, so it is a body change, not a contract change.

func (s *sentinelOwnResources) listPolicies(
	ctx context.Context, orgID, emplID uuid.UUID, f agg.InstanceFilter,
) (agg.InstancePage, error) {
	policies, err := s.sentinelService.ListPolicies(ctx, orgID)
	if err != nil {
		return agg.InstancePage{}, err
	}
	out := make([]agg.ResourceInstance, 0, len(policies))
	for _, p := range policies {
		out = append(out, agg.ResourceInstance{ID: p.ID.String(), Label: p.Name})
	}
	return page(out, f), nil
}

func (s *sentinelOwnResources) listRoles(
	ctx context.Context, orgID, emplID uuid.UUID, f agg.InstanceFilter,
) (agg.InstancePage, error) {
	roles, err := s.sentinelService.ListRoles(ctx, orgID)
	if err != nil {
		return agg.InstancePage{}, err
	}
	out := make([]agg.ResourceInstance, 0, len(roles))
	for _, r := range roles {
		out = append(out, agg.ResourceInstance{ID: r.ID.String(), Label: r.Name})
	}
	return page(out, f), nil
}

func (s *sentinelOwnResources) listGroups(
	ctx context.Context, orgID, emplID uuid.UUID, f agg.InstanceFilter,
) (agg.InstancePage, error) {
	groups, err := s.sentinelService.ListGroups(ctx, orgID)
	if err != nil {
		return agg.InstancePage{}, err
	}
	out := make([]agg.ResourceInstance, 0, len(groups))
	for _, g := range groups {
		out = append(out, agg.ResourceInstance{ID: g.ID.String(), Label: g.Name})
	}
	return page(out, f), nil
}

/* ---------------------------------- <***> --------------------------------- */

// page applies an InstanceFilter to an already-loaded slice.
//
// IDs wins over Query when both are set: an id filter is the builder resolving a
// stored policy's scoped instances, and it must return exactly those, not those
// that also happen to match a search box the user left filled in.
func page(all []agg.ResourceInstance, f agg.InstanceFilter) agg.InstancePage {
	filtered := all

	if len(f.IDs) > 0 {
		want := make(map[string]struct{}, len(f.IDs))
		for _, id := range f.IDs {
			want[id] = struct{}{}
		}
		filtered = filtered[:0:0]
		for _, i := range all {
			if _, ok := want[i.ID]; ok {
				filtered = append(filtered, i)
			}
		}
	} else if q := strings.TrimSpace(strings.ToLower(f.Query)); q != "" {
		filtered = filtered[:0:0]
		for _, i := range all {
			if strings.Contains(strings.ToLower(i.Label), q) {
				filtered = append(filtered, i)
			}
		}
	}

	if f.Limit > 0 && len(filtered) > f.Limit {
		// No cursor: these lists are small enough that a truncated page is a
		// signal to narrow the search, not something to paginate through. A
		// real cursor belongs with the SQL-side filter.
		filtered = filtered[:f.Limit]
	}
	return agg.InstancePage{Instances: filtered}
}
