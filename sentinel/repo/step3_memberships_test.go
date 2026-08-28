package sentinel_repo_test

import (
	"context"
	"testing"

	models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"

	"github.com/google/uuid"
)

func newRole(t *testing.T, r sentinel_repo.SentinelRepositoryInterface, org uuid.UUID) *models.Role {
	t.Helper()
	role := &models.Role{OrganizationID: org, Name: "r-" + uuid.NewString()}
	if err := r.CreateRole(context.Background(), role); err != nil {
		t.Fatalf("seed role: %v", err)
	}
	return role
}

func newGroup(t *testing.T, r sentinel_repo.SentinelRepositoryInterface, org uuid.UUID) *models.Group {
	t.Helper()
	g := &models.Group{OrganizationID: org, Name: "g-" + uuid.NewString()}
	if err := r.CreateGroup(context.Background(), g); err != nil {
		t.Fatalf("seed group: %v", err)
	}
	return g
}

func TestRoleMembershipLifecycle(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, role := uuid.New(), newRole(t, r, orgA)

	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); ok {
		t.Fatal("not a member before assigning")
	}
	if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
		t.Fatalf("assign: %v", err)
	}
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); !ok {
		t.Fatal("must be a member after assigning")
	}

	// reads resolve to ENTITIES, not edges — no N+1 for the caller
	roles, err := r.GetRolesOfUser(ctx, orgA, user)
	if err != nil || len(roles) != 1 || roles[0].ID != role.ID {
		t.Fatalf("GetRolesOfUser = %v (err %v), want the one role", roles, err)
	}
	if roles[0].Name != role.Name {
		t.Fatalf("entity not hydrated: name=%q", roles[0].Name)
	}

	users, _ := r.GetUsersOfRole(ctx, orgA, role.ID)
	if len(users) != 1 || users[0] != user {
		t.Fatalf("GetUsersOfRole = %v, want [%s]", users, user)
	}

	// hard delete: no history, and the pair can be assigned again
	if err := r.RemoveRoleFromUser(ctx, orgA, user, role.ID); err != nil {
		t.Fatalf("remove: %v", err)
	}
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); ok {
		t.Fatal("must not be a member after removal")
	}
	if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
		t.Fatalf("re-assign after removal: %v", err)
	}
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); !ok {
		t.Fatal("re-assign must restore membership")
	}
}

func TestGroupMembershipLifecycle(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, group := uuid.New(), newGroup(t, r, orgA)

	if err := r.AssignUserToGroup(ctx, orgA, user, group.ID); err != nil {
		t.Fatalf("assign: %v", err)
	}
	if ok, _ := r.IsUserInGroup(ctx, orgA, user, group.ID); !ok {
		t.Fatal("must be a member")
	}
	groups, err := r.GetGroupsOfUser(ctx, orgA, user)
	if err != nil || len(groups) != 1 || groups[0].ID != group.ID {
		t.Fatalf("GetGroupsOfUser = %v (err %v)", groups, err)
	}
	if err := r.RemoveUserFromGroup(ctx, orgA, user, group.ID); err != nil {
		t.Fatal(err)
	}
	if groups, _ = r.GetGroupsOfUser(ctx, orgA, user); len(groups) != 0 {
		t.Fatalf("removed membership still listed: %v", groups)
	}
}

// The tables carry no unique index (000038 dropped the composite PK), so
// assign must not create a second row for a pair it already has.
func TestAssignIsIdempotentWithoutAUniqueIndex(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, role := uuid.New(), newRole(t, r, orgA)

	for i := 0; i < 3; i++ {
		if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
			t.Fatalf("assign %d: %v", i, err)
		}
	}
	roles, _ := r.GetRolesOfUser(ctx, orgA, user)
	if len(roles) != 1 {
		t.Fatalf("repeated assign duplicated the membership: %d rows", len(roles))
	}
	users, _ := r.GetUsersOfRole(ctx, orgA, role.ID)
	if len(users) != 1 {
		t.Fatalf("repeated assign duplicated the membership: %d users", len(users))
	}
}

// Memberships must never cross tenants.
func TestMembershipsAreOrgScoped(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, role := uuid.New(), newRole(t, r, orgA)

	if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
		t.Fatal(err)
	}
	if roles, _ := r.GetRolesOfUser(ctx, orgB, user); len(roles) != 0 {
		t.Fatalf("orgB saw orgA's membership: %v", roles)
	}
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgB, user, role.ID); ok {
		t.Fatal("IsRoleAssignedToUser leaked across orgs")
	}
	if users, _ := r.GetUsersOfRole(ctx, orgB, role.ID); len(users) != 0 {
		t.Fatal("GetUsersOfRole leaked across orgs")
	}
	if err := r.RemoveRoleFromUser(ctx, orgB, user, role.ID); err != nil {
		t.Fatalf("cross-org removal should no-op: %v", err)
	}
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); !ok {
		t.Fatal("cross-org removal must NOT have deleted the membership")
	}
}

// A user pointing at a soft-deleted container is not in it — otherwise they
// would keep receiving the policies attached to a role that no longer exists.
func TestSoftDeletedContainersDropOutOfMembershipReads(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user, role, group := uuid.New(), newRole(t, r, orgA), newGroup(t, r, orgA)

	if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
		t.Fatal(err)
	}
	if err := r.AssignUserToGroup(ctx, orgA, user, group.ID); err != nil {
		t.Fatal(err)
	}
	if err := r.DeleteRole(ctx, orgA, role.ID); err != nil {
		t.Fatal(err)
	}
	if err := r.DeleteGroup(ctx, orgA, group.ID); err != nil {
		t.Fatal(err)
	}

	if roles, _ := r.GetRolesOfUser(ctx, orgA, user); len(roles) != 0 {
		t.Fatalf("deleted role still returned: %v", roles)
	}
	if groups, _ := r.GetGroupsOfUser(ctx, orgA, user); len(groups) != 0 {
		t.Fatalf("deleted group still returned: %v", groups)
	}
	// the membership row itself survives — nothing cascades from roles/groups
	if ok, _ := r.IsRoleAssignedToUser(ctx, orgA, user, role.ID); !ok {
		t.Fatal("membership row should still exist; only the entity read filters it")
	}
}

// Multiple memberships resolve in one query, in a stable order.
func TestUserWithSeveralRoles(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	user := uuid.New()
	want := map[uuid.UUID]bool{}
	for i := 0; i < 3; i++ {
		role := newRole(t, r, orgA)
		want[role.ID] = true
		if err := r.AssignRoleToUser(ctx, orgA, user, role.ID); err != nil {
			t.Fatal(err)
		}
	}
	roles, err := r.GetRolesOfUser(ctx, orgA, user)
	if err != nil || len(roles) != 3 {
		t.Fatalf("GetRolesOfUser returned %d roles (err %v), want 3", len(roles), err)
	}
	for _, got := range roles {
		if !want[got.ID] {
			t.Fatalf("unexpected role %s", got.ID)
		}
	}
}
