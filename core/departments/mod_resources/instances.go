package orgDepartmentResources

import (
	"context"
	"fmt"
	"strings"

	agg "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// listDepartments implements [agg.InstanceLister]: the org's departments as
// selectable policy targets, so a role can be scoped to "Housekeeping" rather
// than to every department in the organization.
//
// Filtering happens here, in memory, over the service's org-scoped list. An
// organization has tens of departments, not thousands, so that is the right
// trade for now; the signature already carries Query, IDs and Limit, so pushing
// them into SQL later is a body change rather than a contract change.
func (s *orgDepartmentResources) listDepartments(
	ctx context.Context, orgID, emplID uuid.UUID, filter agg.InstanceFilter,
) (agg.InstancePage, error) {
	departments, err := s.departmentService.ListDepartments(ctx, orgID, emplID)
	if err != nil {
		return agg.InstancePage{}, err
	}

	all := make([]agg.ResourceInstance, 0, len(departments))
	for _, d := range departments {
		label := d.Name
		if d.Code != "" {
			label = fmt.Sprintf("%s (%s)", d.Name, d.Code)
		}
		all = append(all, agg.ResourceInstance{ID: d.ID.String(), Label: label})
	}

	return applyFilter(all, filter), nil
}

/* ---------------------------------- <***> --------------------------------- */

// applyFilter narrows an already-loaded slice.
//
// IDs wins over Query when both are set: an id filter is the builder resolving
// the instances a stored policy is scoped to, and it must return exactly those.
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
