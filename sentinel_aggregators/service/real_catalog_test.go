package sentinelAggregatorService_test

import (
	"context"
	"testing"

	orgPropertyResources "sharpnr-engine/core/org_property/mod_resources"
	sentinelOwnResources "sharpnr-engine/sentinel/mod_resources"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"

	"github.com/google/uuid"
)

// TestRealCatalogPassesValidation boots the aggregator with the PRODUCTION
// providers. Every check in NewSentinelAggregatorService panics, so this failing
// means the real app would refuse to start — which is the point: it catches an
// action-id collision between two shipped modules before deploy rather than
// after, when policies already reference the ambiguous id.
//
// The services are nil because Resources() only returns static descriptors; the
// listers are never invoked here.
func TestRealCatalogPassesValidation(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		orgPropertyResources.NewOrgPropertyResources(nil),
		sentinelOwnResources.NewSentinelOwnResources(nil),
	)

	resources, err := svc.ListResources(context.Background(), uuid.New(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{
		"Property": false, "SentinelPolicy": false, "SentinelRole": false, "SentinelGroup": false,
	}
	for _, r := range resources {
		if _, ok := want[r.Type]; !ok {
			t.Errorf("unexpected resource type in the real catalog: %q", r.Type)
			continue
		}
		want[r.Type] = true
		if r.Label == "" {
			t.Errorf("resource %q has no Label — the picker would render blank", r.Type)
		}
	}
	for rtype, found := range want {
		if !found {
			t.Errorf("resource type missing from the real catalog: %q", rtype)
		}
	}
}

// Every real action id must be uniquely owned across the shipped modules.
//
// This reads the providers' descriptors directly rather than going through
// GetResourceDetails, which would invoke the instance listers and need real
// services behind them. Descriptors are static, so no service is required.
func TestRealCatalogActionNamesAreUnique(t *testing.T) {
	providers := []sentinelAggregatorContracts.ResourceProvider{
		orgPropertyResources.NewOrgPropertyResources(nil),
		sentinelOwnResources.NewSentinelOwnResources(nil),
	}

	owner := map[string]string{}
	total, types := 0, 0
	for _, p := range providers {
		for _, reg := range p.Resources() {
			types++
			for _, a := range reg.Descriptor.Actions {
				total++
				if prev, dup := owner[a.Name]; dup {
					t.Errorf("action %q owned by both %q and %q — a policy scoped to it would reach both",
						a.Name, prev, reg.Descriptor.Type)
				}
				owner[a.Name] = reg.Descriptor.Type
				if a.Label == "" {
					t.Errorf("action %q on %q has no Label", a.Name, reg.Descriptor.Type)
				}
			}
		}
	}
	if total == 0 {
		t.Fatal("the real catalog exposes no actions at all")
	}
	t.Logf("real catalog: %d resource types, %d uniquely-named actions", types, total)
}
