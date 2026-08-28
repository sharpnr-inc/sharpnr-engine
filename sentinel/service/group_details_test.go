package sentinelservice_test

import (
	"context"
	"os"
	"testing"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinelservice "sharpnr-engine/sentinel/service"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newService(t *testing.T) (sentinelservice.SentinelServiceInterface, sentinel_repo.SentinelRepositoryInterface) {
	t.Helper()
	dsn := os.Getenv("SENTINEL_TEST_DSN")
	if dsn == "" {
		t.Skip("SENTINEL_TEST_DSN not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	repo := sentinel_repo.NewSentinelRepo(db)
	return sentinelservice.NewSentinelService(repo, dbtx.NewGormTxManager(db), nil), repo
}

func TestGetGroupDetails(t *testing.T) {
	svc, _ := newService(t)
	ctx := context.Background()
	org := uuid.New()

	groupID, err := svc.CreateGroup(ctx, org, sentinelDtos.CreateGroupsRequestDto{Name: "Finance", Description: "finance team"})
	if err != nil {
		t.Fatalf("create group: %v", err)
	}

	detail, err := svc.GetGroupDetails(ctx, org, groupID)
	if err != nil {
		t.Fatalf("empty group must still resolve: %v", err)
	}
	if detail.Group.Name != "Finance" || detail.Group.Description != "finance team" {
		t.Fatalf("group not hydrated: %+v", detail.Group)
	}
	if detail.MemberIDs == nil || len(detail.MemberIDs) != 0 {
		t.Fatalf("empty group must give an empty slice not nil, got %#v", detail.MemberIDs)
	}
	if len(detail.Policies) != 0 {
		t.Fatalf("want no policies, got %d", len(detail.Policies))
	}
	if detail.CreatedAt.IsZero() {
		t.Error("CreatedAt not populated")
	}
}

func TestGetGroupDetailsWithMembersAndPolicies(t *testing.T) {
	svc, _ := newService(t)
	ctx := context.Background()
	org := uuid.New()

	groupID, err := svc.CreateGroup(ctx, org, sentinelDtos.CreateGroupsRequestDto{Name: "Ops"})
	if err != nil {
		t.Fatal(err)
	}

	pEnabled, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "ops-read", Policy: "permit(principal,action,resource);", IsEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	pDisabled, err := svc.CreatePolicy(ctx, org, sentinelDtos.AddPolicyRequestDto{
		Name: "ops-off", Policy: "permit(principal,action,resource);", IsEnabled: false})
	if err != nil {
		t.Fatal(err)
	}

	_, db := newService(t)
	if err := db.AttachPolicyToGroup(ctx, org, groupID, pEnabled); err != nil {
		t.Fatal(err)
	}
	if err := db.AttachPolicyToGroup(ctx, org, groupID, pDisabled); err != nil {
		t.Fatal(err)
	}
	u1, u2 := uuid.New(), uuid.New()
	if err := db.AssignUserToGroup(ctx, org, u1, groupID); err != nil {
		t.Fatal(err)
	}
	if err := db.AssignUserToGroup(ctx, org, u2, groupID); err != nil {
		t.Fatal(err)
	}

	detail, err := svc.GetGroupDetails(ctx, org, groupID)
	if err != nil {
		t.Fatal(err)
	}
	if len(detail.MemberIDs) != 2 {
		t.Fatalf("want 2 members, got %d", len(detail.MemberIDs))
	}
	// a disabled policy is still REGISTERED, so an admin must see it
	if len(detail.Policies) != 2 {
		t.Fatalf("want both policies (incl. disabled), got %d", len(detail.Policies))
	}
	var sawDisabled bool
	for _, p := range detail.Policies {
		if p.ID == pDisabled {
			sawDisabled = true
			if p.Enabled {
				t.Error("disabled policy reported as enabled")
			}
		}
	}
	if !sawDisabled {
		t.Error("disabled-but-registered policy missing from group details")
	}
}

// Another tenant's group must be not-found, not readable.
func TestGetGroupDetailsIsOrgScoped(t *testing.T) {
	svc, _ := newService(t)
	ctx := context.Background()
	orgA, orgB := uuid.New(), uuid.New()

	groupID, err := svc.CreateGroup(ctx, orgA, sentinelDtos.CreateGroupsRequestDto{Name: "Private"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := svc.GetGroupDetails(ctx, orgB, groupID); err == nil {
		t.Fatal("LEAK: another org read the group's details")
	}
}

func TestGetGroupDetailsUnknownGroup(t *testing.T) {
	svc, _ := newService(t)
	if _, err := svc.GetGroupDetails(context.Background(), uuid.New(), uuid.New()); err == nil {
		t.Fatal("unknown group must error")
	}
}
