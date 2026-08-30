package orgPropertyResources

import (
	"context"

	agg "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// propertyFacts implements [agg.FactResolver]: who owns these properties, and
// what are the attribute values a policy may condition on.
//
// This is the function that makes `resource.city == "Mumbai"` actually work, and
// it is also what populates resource.orgId — the field the engine uses to deny a
// cross-tenant request before any policy runs. Both were previously left to the
// call site and forgotten at every one of them.
//
// Batched because it runs on the authorization path: a list endpoint checking
// fifty properties must not fire fifty queries.
//
// Ids that are absent from the result — deleted, or belonging to another org —
// are absent from the map too. The Guard treats that as "no facts", which denies
// anything conditioned on them rather than guessing.
func (s *orgPropertyResources) propertyFacts(
	ctx context.Context, orgID uuid.UUID, instanceIDs []string,
) (map[string]agg.ResourceFacts, error) {
	ids := make([]uuid.UUID, 0, len(instanceIDs))
	for _, raw := range instanceIDs {
		id, err := uuid.Parse(raw)
		if err != nil {
			// A malformed id cannot name a real property. Skipping it yields no
			// facts for it, which denies — parsing strictly here would turn a
			// bad request into a 500.
			continue
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return map[string]agg.ResourceFacts{}, nil
	}

	rows, err := s.propertyService.GetPropertyFacts(ctx, orgID, ids)
	if err != nil {
		return nil, err
	}

	facts := make(map[string]agg.ResourceFacts, len(rows))
	for _, p := range rows {
		facts[p.ID.String()] = agg.ResourceFacts{
			OrgID: p.OrganizationID,
			Attributes: map[string]any{
				AttrCode:    p.Code,
				AttrName:    p.Name,
				AttrCity:    p.City,
				AttrState:   p.State,
				AttrCountry: p.Country,
			},
		}
	}
	return facts, nil
}

/* ---------------------------------- <***> --------------------------------- */

// attributeValues builds the [agg.ValueResolver] for one attribute: the distinct
// values it takes across the org, so the builder renders a picker rather than a
// text box. Read when the catalog is fetched, never on the decision path.
func (s *orgPropertyResources) attributeValues(attribute string) agg.ValueResolver {
	return func(ctx context.Context, orgID uuid.UUID) ([]agg.AttributeValue, error) {
		raw, err := s.propertyService.GetPropertyAttributeValues(ctx, orgID, attribute)
		if err != nil {
			return nil, err
		}
		values := make([]agg.AttributeValue, 0, len(raw))
		for _, v := range raw {
			values = append(values, agg.Value(v))
		}
		return values, nil
	}
}
