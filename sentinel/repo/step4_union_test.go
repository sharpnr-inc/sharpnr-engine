package sentinel_repo_test

import (
	"context"
	"testing"

	models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// newStore reaches the repo through the engine's narrow read interface rather
// than the authoring one. GetAllPoliciesOfUser is deliberately absent from
// SentinelRepositoryInterface: the service never needs it, and the engine
// depends only on PolicyStore. The assertion below is itself the check that the
// repo still doubles as that read surface.
func newStore(t *testing.T, r sentinel_repo.SentinelRepositoryInterface) sentinelEngine.PolicyStore {
	t.Helper()
	store, ok := r.(sentinelEngine.PolicyStore)
	if !ok {
		t.Fatal("the repo must satisfy sentinelEngine.PolicyStore")
	}
	return store
}

// namedPolicy seeds a policy whose Cedar text embeds its label, so a failure
// message says WHICH policy leaked rather than printing uuids.
func namedPolicy(t *testing.T, r sentinel_repo.SentinelRepositoryInterface, org uuid.UUID, label string, enabled bool) *models.SentinelPolicy {
	t.Helper()
	p := &models.SentinelPolicy{
		OrganizationID: org,
		Name:           label + "-" + uuid.NewString(),
		Policy:         "// " + label + "\npermit(principal,action,resource);",
		Enabled:        enabled,
	}
	if err := r.CreatePolicy(context.Background(), p); err != nil {
		t.Fatalf("seed %s: %v", label, err)
	}
	return p
}

// TestUnionReachesExactlyTheRightPolicies is the load-bearing test of the whole
// subsystem: a mistake here is either a breach (too many) or a lockout (too few).
func TestUnionReachesExactlyTheRightPolicies(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user := uuid.New()

	// A FRESH org, not the shared orgA. The org register reaches every user in
	// its org, so org-wide policies left behind by any other test would
	// legitimately show up here and break the exact-count assertion below.
	// The count is the point of this test, so the org has to be isolated.
	orgA, orgB := uuid.New(), uuid.New()
	_ = orgB

	roleIn := newRole(t, r, orgA)     // user holds this
	roleOut := newRole(t, r, orgA)    // user does NOT hold this
	roleGone := newRole(t, r, orgA)   // held, but soft-deleted
	groupIn := newGroup(t, r, orgA)   // user belongs
	groupGone := newGroup(t, r, orgA) // belongs, but soft-deleted

	for _, m := range []struct {
		assign func() error
	}{
		{func() error { return r.AssignRoleToUser(ctx, orgA, user, roleIn.ID) }},
		{func() error { return r.AssignRoleToUser(ctx, orgA, user, roleGone.ID) }},
		{func() error { return r.AssignUserToGroup(ctx, orgA, user, groupIn.ID) }},
		{func() error { return r.AssignUserToGroup(ctx, orgA, user, groupGone.ID) }},
	} {
		if err := m.assign(); err != nil {
			t.Fatal(err)
		}
	}

	// ---- should reach the user -------------------------------------------
	viaOrg := namedPolicy(t, r, orgA, "via-org", true)
	viaRole := namedPolicy(t, r, orgA, "via-role", true)
	viaGroup := namedPolicy(t, r, orgA, "via-group", true)
	viaUser := namedPolicy(t, r, orgA, "via-user", true)
	viaBoth := namedPolicy(t, r, orgA, "via-org-and-role", true) // dedup case

	mustAttach(t, r.AttachPolicyToOrg(ctx, orgA, viaOrg.ID))
	mustAttach(t, r.AttachPolicyToRole(ctx, orgA, roleIn.ID, viaRole.ID))
	mustAttach(t, r.AttachPolicyToGroup(ctx, orgA, groupIn.ID, viaGroup.ID))
	mustAttach(t, r.AttachPolicyToUser(ctx, orgA, user, viaUser.ID))
	mustAttach(t, r.AttachPolicyToOrg(ctx, orgA, viaBoth.ID))
	mustAttach(t, r.AttachPolicyToRole(ctx, orgA, roleIn.ID, viaBoth.ID))

	// ---- should NOT reach the user ---------------------------------------
	unattached := namedPolicy(t, r, orgA, "unattached", true)
	otherRole := namedPolicy(t, r, orgA, "role-user-lacks", true)
	disabled := namedPolicy(t, r, orgA, "disabled", false)
	softDeleted := namedPolicy(t, r, orgA, "soft-deleted", true)
	detached := namedPolicy(t, r, orgA, "detached", true)
	deletedRole := namedPolicy(t, r, orgA, "via-deleted-role", true)
	deletedGroup := namedPolicy(t, r, orgA, "via-deleted-group", true)
	otherOrg := namedPolicy(t, r, orgB, "other-org", true)

	mustAttach(t, r.AttachPolicyToRole(ctx, orgA, roleOut.ID, otherRole.ID))
	mustAttach(t, r.AttachPolicyToOrg(ctx, orgA, disabled.ID))
	mustAttach(t, r.AttachPolicyToOrg(ctx, orgA, softDeleted.ID))
	mustAttach(t, r.AttachPolicyToRole(ctx, orgA, roleGone.ID, deletedRole.ID))
	mustAttach(t, r.AttachPolicyToGroup(ctx, orgA, groupGone.ID, deletedGroup.ID))
	mustAttach(t, r.AttachPolicyToOrg(ctx, orgB, otherOrg.ID))

	// attached then detached: the register row is soft-deleted
	mustAttach(t, r.AttachPolicyToUser(ctx, orgA, user, detached.ID))
	mustAttach(t, r.DetachPolicyFromUser(ctx, orgA, user, detached.ID))

	if err := r.DeletePolicy(ctx, orgA, softDeleted.ID); err != nil {
		t.Fatal(err)
	}
	if err := r.DeleteRole(ctx, orgA, roleGone.ID); err != nil {
		t.Fatal(err)
	}
	if err := r.DeleteGroup(ctx, orgA, groupGone.ID); err != nil {
		t.Fatal(err)
	}

	// ---- the union -------------------------------------------------------
	records, err := newStore(t, r).GetAllPoliciesOfUser(ctx, orgA, user)
	if err != nil {
		t.Fatalf("GetAllPoliciesOfUser: %v", err)
	}

	got := map[uuid.UUID]int{}
	for _, rec := range records {
		got[rec.ID]++
	}

	wantReach := map[uuid.UUID]string{
		viaOrg.ID: "via-org", viaRole.ID: "via-role", viaGroup.ID: "via-group",
		viaUser.ID: "via-user", viaBoth.ID: "via-org-and-role",
	}
	wantBlocked := map[uuid.UUID]string{
		unattached.ID: "unattached", otherRole.ID: "role-user-lacks",
		disabled.ID: "disabled", softDeleted.ID: "soft-deleted", detached.ID: "detached",
		deletedRole.ID: "via-deleted-role", deletedGroup.ID: "via-deleted-group",
		otherOrg.ID: "other-org",
	}

	for id, label := range wantReach {
		switch got[id] {
		case 0:
			t.Errorf("LOCKOUT: %q should reach the user but did not", label)
		case 1: // correct
		default:
			t.Errorf("DUPLICATE: %q compiled %d times, want 1", label, got[id])
		}
	}
	for id, label := range wantBlocked {
		if got[id] > 0 {
			t.Errorf("LEAK: %q must NOT reach the user, but did", label)
		}
	}
	if len(records) != len(wantReach) {
		t.Errorf("union returned %d policies, want exactly %d", len(records), len(wantReach))
	}

	// the record must carry the Cedar text the engine compiles
	for _, rec := range records {
		if rec.Text == "" {
			t.Errorf("policy %s came back with empty Cedar text", rec.ID)
		}
	}
}

func mustAttach(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("attach: %v", err)
	}
}

// A user with no org-wide policies, no roles, no groups and no grants gets an
// empty set — and an empty slice, not nil, so the engine compiles an empty
// PolicySet and denies rather than erroring.
func TestUnionForUnprovisionedUserIsEmptyNotNil(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	records, err := newStore(t, r).GetAllPoliciesOfUser(ctx, uuid.New(), uuid.New())
	if err != nil {
		t.Fatalf("unprovisioned user must not error: %v", err)
	}
	if records == nil {
		t.Fatal("must return an empty slice, not nil")
	}
	if len(records) != 0 {
		t.Fatalf("want no policies, got %d", len(records))
	}
}

// The same employee id in a different org must see nothing of the first org.
func TestUnionIsTenantIsolated(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user := uuid.New()
	orgA, orgB := uuid.New(), uuid.New()

	p := namedPolicy(t, r, orgA, "orgA-wide", true)
	mustAttach(t, r.AttachPolicyToOrg(ctx, orgA, p.ID))

	store := newStore(t, r)
	inA, err := store.GetAllPoliciesOfUser(ctx, orgA, user)
	if err != nil || len(inA) == 0 {
		t.Fatalf("user must receive orgA's org-wide policy (err %v)", err)
	}
	inB, err := store.GetAllPoliciesOfUser(ctx, orgB, user)
	if err != nil {
		t.Fatal(err)
	}
	for _, rec := range inB {
		if rec.ID == p.ID {
			t.Fatal("LEAK: orgA's org-wide policy reached the same user under orgB")
		}
	}
}

// Re-enabling a policy restores reach without re-attaching: enabled and
// registered are independent switches.
func TestDisablingIsReversibleWithoutReattaching(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, orgA := uuid.New(), uuid.New()
	p := namedPolicy(t, r, orgA, "toggle", true)
	mustAttach(t, r.AttachPolicyToUser(ctx, orgA, user, p.ID))

	if n := len(reach(t, r, orgA, user, p.ID)); n != 1 {
		t.Fatal("enabled+attached policy should reach the user")
	}
	if err := r.SetPolicyEnabled(ctx, orgA, p.ID, false); err != nil {
		t.Fatal(err)
	}
	if n := len(reach(t, r, orgA, user, p.ID)); n != 0 {
		t.Fatal("disabled policy must not reach the user")
	}
	if err := r.SetPolicyEnabled(ctx, orgA, p.ID, true); err != nil {
		t.Fatal(err)
	}
	if n := len(reach(t, r, orgA, user, p.ID)); n != 1 {
		t.Fatal("re-enabled policy must reach the user again, with no re-attach")
	}
}

func reach(t *testing.T, r sentinel_repo.SentinelRepositoryInterface, org, user, policy uuid.UUID) []uuid.UUID {
	t.Helper()
	records, err := newStore(t, r).GetAllPoliciesOfUser(context.Background(), org, user)
	if err != nil {
		t.Fatal(err)
	}
	var hits []uuid.UUID
	for _, rec := range records {
		if rec.ID == policy {
			hits = append(hits, rec.ID)
		}
	}
	return hits
}
