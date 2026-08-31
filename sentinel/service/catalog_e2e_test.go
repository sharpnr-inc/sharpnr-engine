package sentinelservice_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelOwnResources "sharpnr-engine/sentinel/mod_resources"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"

	"github.com/google/uuid"
)

// The catalog wired to a REAL sentinel service and a real database, so the
// instance listers actually run. The unit tests in sentinel_aggregators use
// stubs; this is the path GET /sentinel/resources/:resourceType takes.
func TestResourceCatalogEndToEnd(t *testing.T) {
	svc, repo := newService(t)
	ctx := context.Background()
	org, empl := uuid.New(), uuid.New()

	catalog := sentinelAggregatorService.NewSentinelAggregatorService(
		sentinelOwnResources.NewSentinelOwnResources(svc),
	)

	// seed one of each so the listers have something to return
	roleID, err := svc.CreateRole(ctx, org, empl, sentinelDtos.CreateRoleInput{Name: "Auditor"})
	if err != nil {
		t.Fatal(err)
	}
	groupID, err := svc.CreateGroup(ctx, org, sentinelDtos.CreateGroupsRequestDto{Name: "Stakeholders"})
	if err != nil {
		t.Fatal(err)
	}
	policyID, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "read-only", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	_ = repo

	/* ---- GET /sentinel/resources ---- */
	list, err := catalog.ListResources(ctx, org, empl)
	if err != nil {
		t.Fatalf("ListResources: %v", err)
	}
	if len(list) != 3 {
		t.Fatalf("want 3 sentinel resource types, got %d", len(list))
	}
	raw, _ := json.Marshal(list)
	t.Logf("GET /sentinel/resources -> %s", raw)

	/* ---- GET /sentinel/resources/:resourceType ---- */
	cases := []struct {
		rtype   string
		wantID  uuid.UUID
		minActs int
	}{
		{sentinelOwnResources.ResourceTypeSentinelRole, roleID, 6},
		{sentinelOwnResources.ResourceTypeSentinelGroup, groupID, 6},
		{sentinelOwnResources.ResourceTypeSentinelPolicy, policyID, 7},
	}
	for _, tc := range cases {
		t.Run(tc.rtype, func(t *testing.T) {
			view, err := catalog.GetResourceDetails(ctx, org, empl, tc.rtype, sentinelAggregatorContracts.InstanceFilter{})
			if err != nil {
				t.Fatalf("GetResourceDetails: %v", err)
			}
			if len(view.Actions) < tc.minActs {
				t.Errorf("want >= %d actions, got %d", tc.minActs, len(view.Actions))
			}
			for _, a := range view.Actions {
				if a.Name == "" || a.Label == "" {
					t.Errorf("action rendered blank in the picker: %+v", a)
				}
			}
			// Sentinel's own types declare NO attributes, and that is the
			// design rather than an omission. The previous version advertised
			// `name` and `enabled` with nothing anywhere to stamp them, so a
			// rule like `resource.name != "Owner"` compiled, stored and
			// displayed as active while matching nothing — Cedar skips a policy
			// whose attribute is missing. The contract now refuses to boot on a
			// declaration with no resolver, so an empty list here is the
			// assertion, not a gap in one.
			if len(view.Attributes) != 0 {
				t.Errorf("%s declares attributes with no fact resolver: %+v", tc.rtype, view.Attributes)
			}

			// the lister ran and returned the seeded row
			if !view.HasInstances {
				t.Fatal("HasInstances should be true — this type has a lister")
			}
			var found bool
			for _, inst := range view.Instances {
				if inst.ID == tc.wantID.String() {
					found = true
					if inst.Label == "" {
						t.Error("instance has no label")
					}
				}
			}
			if !found {
				t.Errorf("seeded instance %s missing from %d instances", tc.wantID, len(view.Instances))
			}

			raw, _ := json.Marshal(view)
			t.Logf("GET /sentinel/resources/%s -> %s", tc.rtype, raw)
		})
	}
}

// Instances must be tenant-scoped: the listers delegate to the sentinel
// service, which scopes by org.
func TestResourceInstancesAreTenantScoped(t *testing.T) {
	svc, _ := newService(t)
	ctx := context.Background()
	orgA, orgB, empl := uuid.New(), uuid.New(), uuid.New()

	catalog := sentinelAggregatorService.NewSentinelAggregatorService(
		sentinelOwnResources.NewSentinelOwnResources(svc),
	)
	if _, err := svc.CreateRole(ctx, orgA, empl, sentinelDtos.CreateRoleInput{Name: "SecretRole"}); err != nil {
		t.Fatal(err)
	}

	view, err := catalog.GetResourceDetails(ctx, orgB, empl, sentinelOwnResources.ResourceTypeSentinelRole, sentinelAggregatorContracts.InstanceFilter{})
	if err != nil {
		t.Fatal(err)
	}
	for _, inst := range view.Instances {
		if inst.Label == "SecretRole" {
			t.Fatal("LEAK: orgA's role appeared in orgB's instance picker")
		}
	}
}

// An unknown type must be a typed error so the controller answers 404.
func TestUnknownResourceTypeEndToEnd(t *testing.T) {
	svc, _ := newService(t)
	catalog := sentinelAggregatorService.NewSentinelAggregatorService(
		sentinelOwnResources.NewSentinelOwnResources(svc),
	)
	_, err := catalog.GetResourceDetails(context.Background(), uuid.New(), uuid.New(), "NotAThing", sentinelAggregatorContracts.InstanceFilter{})
	if !errors.Is(err, sentinelAggregatorService.ErrUnknownResourceType) {
		t.Fatalf("want ErrUnknownResourceType (-> 404), got %v", err)
	}
}
