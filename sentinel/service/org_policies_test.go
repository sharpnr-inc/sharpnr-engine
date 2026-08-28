package sentinelservice_test

import (
	"context"
	"testing"

	sentinelDtos "sharpnr-engine/sentinel/dtos"

	"github.com/google/uuid"
)

func hasPolicy(views []sentinelDtos.PolicyView, id uuid.UUID) bool {
	for _, v := range views {
		if v.ID == id {
			return true
		}
	}
	return false
}

// ListOrgPolicies reads the ORG REGISTER; ListPolicies reads the CATALOG.
// A policy the org owns but has registered nowhere reaches nobody, and must not
// appear in the org list — that difference is the point of the split.
func TestListOrgPoliciesReadsTheRegisterNotTheCatalog(t *testing.T) {
	svc, repo := newService(t)
	ctx := context.Background()
	org := uuid.New()

	orgWide, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "org-wide", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	unattached, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "unattached", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	roleOnly, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "role-only", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}

	roleID := uuid.New()
	if err := repo.AttachPolicyToRole(ctx, org, roleID, roleOnly); err != nil {
		t.Fatal(err)
	}
	if err := repo.AttachPolicyToOrg(ctx, org, orgWide); err != nil {
		t.Fatal(err)
	}

	// the register: only what reaches everyone
	orgList, err := svc.ListOrgPolicies(ctx, org)
	if err != nil {
		t.Fatal(err)
	}
	if !hasPolicy(orgList, orgWide) {
		t.Error("the org-registered policy is missing from ListOrgPolicies")
	}
	if hasPolicy(orgList, unattached) {
		t.Error("an UNATTACHED policy appeared in ListOrgPolicies — it reaches nobody")
	}
	if hasPolicy(orgList, roleOnly) {
		t.Error("a ROLE-registered policy appeared in ListOrgPolicies — it reaches only that role")
	}
	if len(orgList) != 1 {
		t.Errorf("ListOrgPolicies returned %d, want exactly 1", len(orgList))
	}

	// the catalog: everything the org owns
	catalog, err := svc.ListPolicies(ctx, org)
	if err != nil {
		t.Fatal(err)
	}
	if len(catalog) != 3 {
		t.Errorf("ListPolicies (catalog) returned %d, want all 3", len(catalog))
	}
}

// Detaching removes it from the org list without touching the policy row.
func TestListOrgPoliciesFollowsDetach(t *testing.T) {
	svc, repo := newService(t)
	ctx := context.Background()
	org := uuid.New()

	id, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "toggle-reach", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	if err := repo.AttachPolicyToOrg(ctx, org, id); err != nil {
		t.Fatal(err)
	}
	if list, _ := svc.ListOrgPolicies(ctx, org); !hasPolicy(list, id) {
		t.Fatal("should be listed after attach")
	}
	if err := repo.DetachPolicyFromOrg(ctx, org, id); err != nil {
		t.Fatal(err)
	}
	if list, _ := svc.ListOrgPolicies(ctx, org); hasPolicy(list, id) {
		t.Fatal("should NOT be listed after detach")
	}
	if _, err := svc.GetPolicy(ctx, org, id); err != nil {
		t.Fatalf("detach must leave the policy row intact: %v", err)
	}
}

func TestListOrgPoliciesIsTenantIsolated(t *testing.T) {
	svc, repo := newService(t)
	ctx := context.Background()
	orgA, orgB := uuid.New(), uuid.New()

	id, err := svc.CreatePolicy(ctx, orgA, sentinelDtos.AddPolicyRequestDto{
		Name: "a-only", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	if err := repo.AttachPolicyToOrg(ctx, orgA, id); err != nil {
		t.Fatal(err)
	}
	if list, _ := svc.ListOrgPolicies(ctx, orgB); len(list) != 0 {
		t.Fatalf("LEAK: orgB saw %d of orgA's org-wide policies", len(list))
	}
}
