package sentinelengine

import (
	"context"
	"testing"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// This test doubles as a runnable, copy-pasteable EXAMPLE of how to use the
// engine. It uses an in-memory PolicyStore so it needs no database. In the real
// app you implement the same PolicyStore interface against sentinel_engine/repo.

/* ----------------------- example in-memory PolicyStore -------------------- */

// memStore is the simplest possible PolicyStore: it returns the same policies
// and memberships for everyone. Swap this for your DB-backed repo, where
// PoliciesFor/MembershipsOf are driven by the principal's actual roles/groups.
type memStore struct {
	policies    []sentinelengine_models.Policy
	memberships []Membership
}

func (m memStore) PoliciesFor(_ context.Context, _, _ string) ([]sentinelengine_models.Policy, error) {
	return m.policies, nil
}

func (m memStore) MembershipsOf(_ context.Context, _, _ string) ([]Membership, error) {
	return m.memberships, nil
}

/* --------------------------------- tests --------------------------------- */

// Allow: alice is in the "engineer" role; engineers may read Documents.
func TestEvaluate_RolePolicyAllows(t *testing.T) {
	store := memStore{
		policies: []sentinelengine_models.Policy{{
			ID:            "engineer-read-docs",
			Effect:        sentinelengine_models.Permit,
			PrincipalType: sentinelengine_models.PrincipalRole,
			PrincipalID:   "engineer",
			Actions:       []string{"read", "list"},
			ResourceType:  "Document",
		}},
		memberships: []Membership{{Type: "Role", ID: "engineer"}},
	}
	sentinel := New(store)

	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: "alice",
		Action:       "read",
		ResourceType: "Document", ResourceID: "report-q3",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Allowed {
		t.Fatalf("expected allowed, got denied: %s", got.Reason)
	}
}

// Default deny: there IS a read policy, but the requested action ("delete")
// isn't covered, so nothing permits it.
func TestEvaluate_NoMatchingPolicyDenies(t *testing.T) {
	store := memStore{
		policies: []sentinelengine_models.Policy{{
			ID:            "engineer-read-docs",
			Effect:        sentinelengine_models.Permit,
			PrincipalType: sentinelengine_models.PrincipalRole,
			PrincipalID:   "engineer",
			Actions:       []string{"read"},
			ResourceType:  "Document",
		}},
		memberships: []Membership{{Type: "Role", ID: "engineer"}},
	}
	sentinel := New(store)

	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: "alice",
		Action:       "delete",
		ResourceType: "Document", ResourceID: "report-q3",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Allowed {
		t.Fatal("expected denied for an un-permitted action")
	}
}

// Forbid wins: a permit and a forbid both match — Cedar denies regardless.
func TestEvaluate_ForbidOverridesPermit(t *testing.T) {
	store := memStore{
		policies: []sentinelengine_models.Policy{
			{
				ID:            "engineer-read-docs",
				Effect:        sentinelengine_models.Permit,
				PrincipalType: sentinelengine_models.PrincipalRole,
				PrincipalID:   "engineer",
				Actions:       []string{"read"},
				ResourceType:  "Document",
			},
			{
				ID:            "contractors-cannot-read-docs",
				Effect:        sentinelengine_models.Forbid,
				PrincipalType: sentinelengine_models.PrincipalGroup,
				PrincipalID:   "contractors",
				Actions:       []string{"read"},
				ResourceType:  "Document",
			},
		},
		// alice is both an engineer (permit) and a contractor (forbid).
		memberships: []Membership{
			{Type: "Role", ID: "engineer"},
			{Type: "Group", ID: "contractors"},
		},
	}
	sentinel := New(store)

	got, err := sentinel.Evaluate(context.Background(), EvaluationRequest{
		PrincipalType: "User", PrincipalID: "alice",
		Action:       "read",
		ResourceType: "Document", ResourceID: "report-q3",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Allowed {
		t.Fatal("expected denied: a forbid policy must override the permit")
	}
}

// ABAC via RawCedar + context: only allow when context.mfa == true.
func TestEvaluate_ContextDrivenPolicy(t *testing.T) {
	store := memStore{
		policies: []sentinelengine_models.Policy{{
			ID: "read-docs-with-mfa",
			RawCedar: `permit (
			  principal in Role::"engineer",
			  action == Action::"read",
			  resource is Document
			) when { context.mfa == true };`,
		}},
		memberships: []Membership{{Type: "Role", ID: "engineer"}},
	}
	sentinel := New(store)

	base := EvaluationRequest{
		PrincipalType: "User", PrincipalID: "alice",
		Action:       "read",
		ResourceType: "Document", ResourceID: "report-q3",
	}

	// Without MFA -> denied.
	withoutMFA := base
	withoutMFA.Context = map[string]any{"mfa": false}
	if got, _ := sentinel.Evaluate(context.Background(), withoutMFA); got.Allowed {
		t.Fatal("expected denied when context.mfa == false")
	}

	// With MFA -> allowed.
	withMFA := base
	withMFA.Context = map[string]any{"mfa": true}
	if got, _ := sentinel.Evaluate(context.Background(), withMFA); !got.Allowed {
		t.Fatal("expected allowed when context.mfa == true")
	}
}
