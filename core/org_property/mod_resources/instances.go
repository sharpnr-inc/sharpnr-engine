package orgPropertyResources

import (
	"context"
	"fmt"
	"strings"

	agg "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// listProperties implements [agg.InstanceLister]: the org's properties as
// selectable policy targets.
//
// Property is the type the instance picker matters most for — most other
// resources are reached by scoping to their property rather than being picked
// individually.
//
// Filtering happens here, in memory, over the service's existing org-scoped
// list. That is the wrong shape once an org has thousands of properties; the
// signature already carries Query, IDs and Limit, so pushing them into SQL is a
// body change rather than a contract change.
//
// ⚠ emplID is accepted and not yet used to narrow the result. Until it is, this
// shows every property name in the org to anyone who opens the policy builder —
// exactly the screen a property-scoped user should not be able to read the whole
// estate out of.
func (s *orgPropertyResources) listProperties(
	ctx context.Context, orgID, emplID uuid.UUID, filter agg.InstanceFilter,
) (agg.InstancePage, error) {
	properties, err := s.propertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return agg.InstancePage{}, err
	}

	all := make([]agg.ResourceInstance, 0, len(properties))
	for _, p := range properties {
		label := p.Name
		if p.Code != "" {
			label = fmt.Sprintf("%s (%s)", p.Name, p.Code)
		}
		all = append(all, agg.ResourceInstance{ID: p.ID.String(), Label: label})
	}
	return applyFilter(all, filter), nil
}

/* ---------------------------------- <***> --------------------------------- */

// applyFilter narrows an already-loaded slice.
//
// IDs wins over Query when both are set: an id filter is the builder resolving
// the instances a stored policy is scoped to, and it must return exactly those —
// not those that also happen to match a search box the user left filled in.
func applyFilter(all []agg.ResourceInstance, f agg.InstanceFilter) agg.InstancePage {
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
		filtered = filtered[:f.Limit]
	}
	return agg.InstancePage{Instances: filtered}
}
