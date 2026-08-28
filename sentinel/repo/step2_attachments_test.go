package sentinel_repo_test

import (
	"context"
	"testing"

	models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"

	"github.com/google/uuid"
)

func newPolicy(t *testing.T, r sentinel_repo.SentinelRepositoryInterface, org uuid.UUID) uuid.UUID {
	t.Helper()
	p := &models.SentinelPolicy{OrganizationID: org, Name: "p-" + uuid.NewString(), Policy: "permit(principal,action,resource);", Enabled: true}
	if err := r.CreatePolicy(context.Background(), p); err != nil {
		t.Fatalf("seed policy: %v", err)
	}
	return p.ID
}

func has(policies []*models.SentinelPolicy, id uuid.UUID) bool { return count(policies, id) > 0 }

// count is what the duplicate check needs: ListPoliciesOfOrg returns every
// org-wide policy (the org IS the principal, so there is one bucket per org),
// so total length says nothing about whether one registration was duplicated.
func count(policies []*models.SentinelPolicy, id uuid.UUID) int {
	n := 0
	for _, p := range policies {
		if p.ID == id {
			n++
		}
	}
	return n
}

// The core invariant: a policy reaches a principal iff a live row registers it.
func TestAttachDetachAcrossAllFourRegisters(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	cases := []struct {
		kind   string
		attach func(principal, policy uuid.UUID) error
		detach func(principal, policy uuid.UUID) error
		list   func(principal uuid.UUID) ([]*models.SentinelPolicy, error)
		is     func(principal, policy uuid.UUID) (bool, error)
	}{
		{"org",
			func(_, pol uuid.UUID) error { return r.AttachPolicyToOrg(ctx, orgA, pol) },
			func(_, pol uuid.UUID) error { return r.DetachPolicyFromOrg(ctx, orgA, pol) },
			func(_ uuid.UUID) ([]*models.SentinelPolicy, error) { return r.ListPoliciesOfOrg(ctx, orgA) },
			func(_, pol uuid.UUID) (bool, error) { return r.IsPolicyAttachedToOrg(ctx, orgA, pol) }},
		{"role",
			func(pr, pol uuid.UUID) error { return r.AttachPolicyToRole(ctx, orgA, pr, pol) },
			func(pr, pol uuid.UUID) error { return r.DetachPolicyFromRole(ctx, orgA, pr, pol) },
			func(pr uuid.UUID) ([]*models.SentinelPolicy, error) { return r.ListPoliciesOfRole(ctx, orgA, pr) },
			func(pr, pol uuid.UUID) (bool, error) { return r.IsPolicyAttachedToRole(ctx, orgA, pr, pol) }},
		{"group",
			func(pr, pol uuid.UUID) error { return r.AttachPolicyToGroup(ctx, orgA, pr, pol) },
			func(pr, pol uuid.UUID) error { return r.DetachPolicyFromGroup(ctx, orgA, pr, pol) },
			func(pr uuid.UUID) ([]*models.SentinelPolicy, error) { return r.ListPoliciesOfGroup(ctx, orgA, pr) },
			func(pr, pol uuid.UUID) (bool, error) { return r.IsPolicyAttachedToGroup(ctx, orgA, pr, pol) }},
		{"user",
			func(pr, pol uuid.UUID) error { return r.AttachPolicyToUser(ctx, orgA, pr, pol) },
			func(pr, pol uuid.UUID) error { return r.DetachPolicyFromUser(ctx, orgA, pr, pol) },
			func(pr uuid.UUID) ([]*models.SentinelPolicy, error) { return r.ListPoliciesOfUser(ctx, orgA, pr) },
			func(pr, pol uuid.UUID) (bool, error) { return r.IsPolicyAttachedToUser(ctx, orgA, pr, pol) }},
	}

	for _, tc := range cases {
		t.Run(tc.kind, func(t *testing.T) {
			principal, policy := uuid.New(), newPolicy(t, r, orgA)

			if ok, _ := tc.is(principal, policy); ok {
				t.Fatal("must not be attached before attaching")
			}
			if err := tc.attach(principal, policy); err != nil {
				t.Fatalf("attach: %v", err)
			}
			if ok, _ := tc.is(principal, policy); !ok {
				t.Fatal("must be attached after attaching")
			}
			got, err := tc.list(principal)
			if err != nil || !has(got, policy) {
				t.Fatalf("list must contain the policy (err %v)", err)
			}

			// idempotent: attaching a live pair again must not error or duplicate
			if err := tc.attach(principal, policy); err != nil {
				t.Fatalf("re-attach must be a no-op, got: %v", err)
			}
			if got, _ = tc.list(principal); count(got, policy) != 1 {
				t.Fatalf("re-attach duplicated the registration: policy appears %d times", count(got, policy))
			}

			// detach is a soft delete: the policy row survives, the reach stops
			if err := tc.detach(principal, policy); err != nil {
				t.Fatalf("detach: %v", err)
			}
			if ok, _ := tc.is(principal, policy); ok {
				t.Fatal("must not be attached after detaching")
			}
			if got, _ = tc.list(principal); has(got, policy) {
				t.Fatal("detached policy must not be listed")
			}
			if _, err := r.GetPolicy(ctx, orgA, policy); err != nil {
				t.Fatalf("detach must leave the policy row untouched: %v", err)
			}

			// re-attach after a soft delete must work (partial unique index)
			if err := tc.attach(principal, policy); err != nil {
				t.Fatalf("re-attach after detach must succeed, got: %v", err)
			}
			if ok, _ := tc.is(principal, policy); !ok {
				t.Fatal("re-attached policy must reach the principal again")
			}
		})
	}
}

// Registrations must never reach across tenants.
func TestAttachmentsAreOrgScoped(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	role, policy := uuid.New(), newPolicy(t, r, orgA)

	if err := r.AttachPolicyToRole(ctx, orgA, role, policy); err != nil {
		t.Fatal(err)
	}
	got, _ := r.ListPoliciesOfRole(ctx, orgB, role)
	if len(got) != 0 {
		t.Fatalf("orgB must not see orgA's registration, got %d", len(got))
	}
	if ok, _ := r.IsPolicyAttachedToRole(ctx, orgB, role, policy); ok {
		t.Fatal("IsPolicyAttachedToRole leaked across orgs")
	}
	if err := r.DetachPolicyFromRole(ctx, orgB, role, policy); err != nil {
		t.Fatalf("cross-org detach should no-op: %v", err)
	}
	if ok, _ := r.IsPolicyAttachedToRole(ctx, orgA, role, policy); !ok {
		t.Fatal("cross-org detach must NOT have removed the registration")
	}
}

// The org register is the degenerate case: principal column IS organization_id.
func TestOrgRegisterDoesNotBleedBetweenOrgs(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	pA, pB := newPolicy(t, r, orgA), newPolicy(t, r, orgB)
	if err := r.AttachPolicyToOrg(ctx, orgA, pA); err != nil {
		t.Fatal(err)
	}
	if err := r.AttachPolicyToOrg(ctx, orgB, pB); err != nil {
		t.Fatal(err)
	}

	a, _ := r.ListPoliciesOfOrg(ctx, orgA)
	if !has(a, pA) || has(a, pB) {
		t.Fatal("ListPoliciesOfOrg(orgA) must contain only orgA's org-wide policy")
	}
	b, _ := r.ListPoliciesOfOrg(ctx, orgB)
	if !has(b, pB) || has(b, pA) {
		t.Fatal("ListPoliciesOfOrg(orgB) must contain only orgB's org-wide policy")
	}
}

// The reverse lookup that replaces the removed `scope` label.
func TestListPolicyAttachments(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	policy := newPolicy(t, r, orgA)
	role, group := uuid.New(), uuid.New()

	// unattached is visibly unattached — the state `scope` could never show
	att, err := r.ListPolicyAttachments(ctx, orgA, policy)
	if err != nil {
		t.Fatal(err)
	}
	if len(att) != 0 {
		t.Fatalf("a policy attached to nothing must report no attachments, got %v", att)
	}

	if err := r.AttachPolicyToRole(ctx, orgA, role, policy); err != nil {
		t.Fatal(err)
	}
	if err := r.AttachPolicyToGroup(ctx, orgA, group, policy); err != nil {
		t.Fatal(err)
	}
	if err := r.AttachPolicyToOrg(ctx, orgA, policy); err != nil {
		t.Fatal(err)
	}

	att, err = r.ListPolicyAttachments(ctx, orgA, policy)
	if err != nil {
		t.Fatal(err)
	}
	if len(att[sentinel_repo.PrincipalRole]) != 1 || att[sentinel_repo.PrincipalRole][0] != role {
		t.Fatalf("role attachment wrong: %v", att[sentinel_repo.PrincipalRole])
	}
	if len(att[sentinel_repo.PrincipalGroup]) != 1 || att[sentinel_repo.PrincipalGroup][0] != group {
		t.Fatalf("group attachment wrong: %v", att[sentinel_repo.PrincipalGroup])
	}
	if len(att[sentinel_repo.PrincipalOrg]) != 1 {
		t.Fatalf("org attachment wrong: %v", att[sentinel_repo.PrincipalOrg])
	}
	if _, ok := att[sentinel_repo.PrincipalUser]; ok {
		t.Fatal("no user grant was made; the kind must be absent, not empty")
	}
}

// Disabled policies still appear on admin screens — enabled is the engine's
// filter, not the register's.
func TestListIncludesDisabledPolicies(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	role, policy := uuid.New(), newPolicy(t, r, orgA)

	if err := r.AttachPolicyToRole(ctx, orgA, role, policy); err != nil {
		t.Fatal(err)
	}
	if err := r.SetPolicyEnabled(ctx, orgA, policy, false); err != nil {
		t.Fatal(err)
	}
	got, _ := r.ListPoliciesOfRole(ctx, orgA, role)
	if !has(got, policy) {
		t.Fatal("a disabled policy must still be listed as attached")
	}
}

// A soft-deleted policy must drop out of every register read.
func TestSoftDeletedPolicyDropsOutOfRegisters(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	role, policy := uuid.New(), newPolicy(t, r, orgA)

	if err := r.AttachPolicyToRole(ctx, orgA, role, policy); err != nil {
		t.Fatal(err)
	}
	if err := r.DeletePolicy(ctx, orgA, policy); err != nil {
		t.Fatal(err)
	}
	got, _ := r.ListPoliciesOfRole(ctx, orgA, role)
	if has(got, policy) {
		t.Fatal("a soft-deleted policy must not be returned by the register read")
	}
}
