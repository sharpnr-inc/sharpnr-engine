package sentinelengine

import (
	"context"
	"testing"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// This file is the end-to-end EXAMPLE of the role/permission/employee_role model
// (reusable roles + per-property scope + CRUD). It shows exactly what your real
// repo-backed PolicyStore has to do — just from in-memory maps instead of a DB.

/* ------------------- an example store over the three models --------------- */

type rbacStore struct {
	roles       map[uuid.UUID]sentinelengine_models.Role
	permissions []sentinelengine_models.Permission // all permission rows
	assignments map[uuid.UUID][]uuid.UUID          // employeeID -> roleIDs
}

// PoliciesFor: find the employee's roles, gather each role's permission rows,
// and convert every row into a Cedar policy. This is the whole job.
func (s rbacStore) PoliciesFor(_ context.Context, _, principalID string) ([]sentinelengine_models.Policy, error) {
	empID := uuid.MustParse(principalID)
	roleIDs := s.assignments[empID]

	var policies []sentinelengine_models.Policy
	for _, roleID := range roleIDs {
		for _, perm := range s.permissions {
			if perm.RoleID != roleID {
				continue
			}
			if policy, ok := perm.ToPolicy(roleID); ok {
				policies = append(policies, policy)
			}
		}
	}
	return policies, nil
}

// MembershipsOf: the roles the employee holds become Cedar "parents".
func (s rbacStore) MembershipsOf(_ context.Context, _, principalID string) ([]Membership, error) {
	empID := uuid.MustParse(principalID)
	var out []Membership
	for _, roleID := range s.assignments[empID] {
		out = append(out, Membership{Type: "Role", ID: roleID.String()})
	}
	return out, nil
}

/* --------------------------------- fixture -------------------------------- */

// scenario sets up one org with two properties and two employees:
//   - alice: Super Admin (everything, everywhere in the org)
//   - bob:   Mumbai Manager (view+edit Restaurants at Mumbai only)
type scenario struct {
	store    rbacStore
	mumbaiID uuid.UUID
	puneID   uuid.UUID
	aliceID  uuid.UUID
	bobID    uuid.UUID
}

func newScenario() scenario {
	orgID := uuid.New()
	mumbai := uuid.New()
	pune := uuid.New()

	superAdminRole := uuid.New()
	mumbaiMgrRole := uuid.New()
	alice := uuid.New()
	bob := uuid.New()

	return scenario{
		mumbaiID: mumbai,
		puneID:   pune,
		aliceID:  alice,
		bobID:    bob,
		store: rbacStore{
			roles: map[uuid.UUID]sentinelengine_models.Role{
				superAdminRole: {ID: superAdminRole, OrganizationID: orgID, Name: "Super Admin", IsSystem: true},
				mumbaiMgrRole:  {ID: mumbaiMgrRole, OrganizationID: orgID, Name: "Mumbai Manager"},
			},
			permissions: []sentinelengine_models.Permission{
				// Super Admin: all types, all properties (PropertyID nil), all CRUD.
				{
					ID: uuid.New(), RoleID: superAdminRole,
					PropertyID:   nil,
					ResourceType: "*",
					CanView:      true, CanCreate: true, CanEdit: true, CanDelete: true,
				},
				// Mumbai Manager: Restaurants at Mumbai only, view + edit.
				{
					ID: uuid.New(), RoleID: mumbaiMgrRole,
					PropertyID:   &mumbai,
					ResourceType: "Restaurant",
					CanView:      true, CanEdit: true,
				},
			},
			assignments: map[uuid.UUID][]uuid.UUID{
				alice: {superAdminRole},
				bob:   {mumbaiMgrRole},
			},
		},
	}
}

/* ---------------------------------- tests --------------------------------- */

func TestRBAC_MumbaiManagerCanEditMumbaiRestaurant(t *testing.T) {
	sc := newScenario()
	sentinel := New(sc.store)

	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: sc.bobID.String(),
		Action:       "edit",
		ResourceType: "Restaurant", ResourceID: uuid.New().String(),
		ResourceParentType: "Property", ResourceParentID: sc.mumbaiID.String(),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Allowed {
		t.Fatalf("bob should edit Mumbai restaurants, got denied: %s", got.Reason)
	}
}

func TestRBAC_MumbaiManagerCannotEditPuneRestaurant(t *testing.T) {
	sc := newScenario()
	sentinel := New(sc.store)

	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: sc.bobID.String(),
		Action:       "edit",
		ResourceType: "Restaurant", ResourceID: uuid.New().String(),
		ResourceParentType: "Property", ResourceParentID: sc.puneID.String(), // different property
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Allowed {
		t.Fatal("bob must NOT touch Pune — his grant is scoped to Mumbai")
	}
}

func TestRBAC_MumbaiManagerCannotDelete(t *testing.T) {
	sc := newScenario()
	sentinel := New(sc.store)

	got, _ := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: sc.bobID.String(),
		Action:       "delete", // only view+edit granted
		ResourceType: "Restaurant", ResourceID: uuid.New().String(),
		ResourceParentType: "Property", ResourceParentID: sc.mumbaiID.String(),
	})

	if got.Allowed {
		t.Fatal("bob has no delete permission")
	}
}

func TestRBAC_SuperAdminCanDoAnythingAnywhere(t *testing.T) {
	sc := newScenario()
	sentinel := New(sc.store)

	// A delete on a Hotel at Pune — a type/property/action bob can't touch,
	// but the Super Admin wildcard covers everything.
	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: sc.aliceID.String(),
		Action:       "delete",
		ResourceType: "Hotel", ResourceID: uuid.New().String(),
		ResourceParentType: "Property", ResourceParentID: sc.puneID.String(),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Allowed {
		t.Fatalf("super admin should be allowed everything, got denied: %s", got.Reason)
	}
}
