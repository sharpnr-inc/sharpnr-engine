package sentinel_repo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	orgA = uuid.MustParse("aaaaaaaa-0000-0000-0000-00000000000a")
	orgB = uuid.MustParse("bbbbbbbb-0000-0000-0000-00000000000b")
)

func newRepo(t *testing.T) sentinel_repo.SentinelRepositoryInterface {
	t.Helper()
	dsn := os.Getenv("SENTINEL_TEST_DSN")
	if dsn == "" {
		t.Skip("SENTINEL_TEST_DSN not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	return sentinel_repo.NewSentinelRepo(db)
}

// Every entity read must be scoped by org: fetching another tenant's row by id
// returns not-found rather than the row.
func TestPolicyIsOrgScoped(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	p := &models.SentinelPolicy{OrganizationID: orgA, Name: "p-" + uuid.NewString(), Policy: "permit(principal,action,resource);", Enabled: true}
	if err := r.CreatePolicy(ctx, p); err != nil {
		t.Fatalf("create: %v", err)
	}
	if p.ID == uuid.Nil {
		t.Fatal("CreatePolicy must generate the id when unset")
	}

	if _, err := r.GetPolicy(ctx, orgA, p.ID); err != nil {
		t.Fatalf("owner org must read its own policy: %v", err)
	}
	if _, err := r.GetPolicy(ctx, orgB, p.ID); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("other org must NOT read it, got err=%v", err)
	}

	// writes are scoped too: orgB must not be able to flip orgA's policy
	if err := r.SetPolicyEnabled(ctx, orgB, p.ID, false); err != nil {
		t.Fatalf("cross-org update should no-op, not error: %v", err)
	}
	got, _ := r.GetPolicy(ctx, orgA, p.ID)
	if !got.Enabled {
		t.Fatal("cross-org SetPolicyEnabled must NOT have taken effect")
	}
	if err := r.SetPolicyEnabled(ctx, orgA, p.ID, false); err != nil {
		t.Fatalf("same-org update: %v", err)
	}
	if got, _ = r.GetPolicy(ctx, orgA, p.ID); got.Enabled {
		t.Fatal("same-org SetPolicyEnabled must take effect")
	}

	// delete is scoped, and is a soft delete
	if err := r.DeletePolicy(ctx, orgB, p.ID); err != nil {
		t.Fatalf("cross-org delete should no-op: %v", err)
	}
	if _, err := r.GetPolicy(ctx, orgA, p.ID); err != nil {
		t.Fatal("cross-org delete must NOT have removed the row")
	}
	if err := r.DeletePolicy(ctx, orgA, p.ID); err != nil {
		t.Fatalf("same-org delete: %v", err)
	}
	if _, err := r.GetPolicy(ctx, orgA, p.ID); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatal("soft-deleted policy must not be returned")
	}
}

func TestRoleIsOrgScoped(t *testing.T) {
	r, ctx := newRepo(t), context.Background()
	name := "role-" + uuid.NewString()

	if err := r.CreateRole(ctx, &models.Role{OrganizationID: orgA, Name: name}); err != nil {
		t.Fatalf("create: %v", err)
	}
	role, err := r.GetRoleByName(ctx, orgA, name)
	if err != nil {
		t.Fatalf("by name: %v", err)
	}

	if _, err := r.GetRole(ctx, orgB, role.ID); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("other org must not read the role, got %v", err)
	}

	// RoleExists must not leak across tenants — the same name is free in orgB
	if ok, _ := r.RoleExists(ctx, orgA, name); !ok {
		t.Fatal("RoleExists false negative in owning org")
	}
	if ok, _ := r.RoleExists(ctx, orgB, name); ok {
		t.Fatal("RoleExists leaked across orgs")
	}

	// partial update: nil fields untouched
	newDesc := "updated"
	if err := r.UpdateRole(ctx, orgA, role.ID, &sentinelDtos.UpdateRoleDetailsRequestDto{Description: &newDesc}); err != nil {
		t.Fatalf("update: %v", err)
	}
	got, _ := r.GetRole(ctx, orgA, role.ID)
	if got.Description != newDesc || got.Name != name {
		t.Fatalf("partial update wrong: name=%q desc=%q", got.Name, got.Description)
	}
	if err := r.UpdateRole(ctx, orgB, role.ID, &sentinelDtos.UpdateRoleDetailsRequestDto{Name: ptr("hijacked")}); err != nil {
		t.Fatalf("cross-org update should no-op: %v", err)
	}
	if got, _ = r.GetRole(ctx, orgA, role.ID); got.Name != name {
		t.Fatal("cross-org UpdateRole must NOT have taken effect")
	}
}

func TestGroupScopingAndNesting(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	parent := &models.Group{OrganizationID: orgA, Name: "parent-" + uuid.NewString(), Description: "p"}
	if err := r.CreateGroup(ctx, parent); err != nil {
		t.Fatalf("create parent: %v", err)
	}
	child := &models.Group{OrganizationID: orgA, Name: "child-" + uuid.NewString(), ParentGroupID: &parent.ID}
	if err := r.CreateGroup(ctx, child); err != nil {
		t.Fatalf("create child: %v", err)
	}

	kids, err := r.ListChildGroups(ctx, orgA, parent.ID)
	if err != nil || len(kids) != 1 || kids[0].ID != child.ID {
		t.Fatalf("ListChildGroups = %v (err %v), want the one child", kids, err)
	}
	if kids, _ = r.ListChildGroups(ctx, orgB, parent.ID); len(kids) != 0 {
		t.Fatal("ListChildGroups leaked across orgs")
	}

	// UpdateGroup writes description (added in migration 000033)
	parent.Description = "renamed"
	if err := r.UpdateGroup(ctx, orgA, parent); err != nil {
		t.Fatalf("update group: %v", err)
	}
	got, _ := r.GetGroup(ctx, orgA, parent.ID)
	if got.Description != "renamed" {
		t.Fatalf("description not persisted, got %q", got.Description)
	}
	if _, err := r.GetGroup(ctx, orgB, parent.ID); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatal("other org must not read the group")
	}
}

// ListX must only ever return the caller's own tenant rows.
func TestListsAreTenantIsolated(t *testing.T) {
	r, ctx := newRepo(t), context.Background()

	if err := r.CreatePolicy(ctx, &models.SentinelPolicy{OrganizationID: orgB, Name: "b-" + uuid.NewString(), Policy: "permit(principal,action,resource);"}); err != nil {
		t.Fatal(err)
	}
	aPolicies, err := r.ListPolicies(ctx, orgA)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range aPolicies {
		if p.OrganizationID != orgA {
			t.Fatalf("ListPolicies(orgA) returned a row from %s", p.OrganizationID)
		}
	}
	roles, _ := r.ListRoles(ctx, orgA)
	for _, x := range roles {
		if x.OrganizationID != orgA {
			t.Fatalf("ListRoles(orgA) returned a row from %s", x.OrganizationID)
		}
	}
	groups, _ := r.ListGroups(ctx, orgA)
	for _, x := range groups {
		if x.OrganizationID != orgA {
			t.Fatalf("ListGroups(orgA) returned a row from %s", x.OrganizationID)
		}
	}
}

func ptr(s string) *string { return &s }
