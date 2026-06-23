package sentinelservice

import (
	"strings"
	"testing"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// The generator's one job is to emit Cedar that (a) compiles and (b) targets the
// right role/actions. These tests pin both: every generated string is fed back
// through the real cedar-go parser, so a malformed template fails loudly here
// instead of denying a whole org closed at runtime.

func mustParse(t *testing.T, text string) {
	t.Helper()
	var p cedar.Policy
	if err := p.UnmarshalCedar([]byte(text)); err != nil {
		t.Fatalf("generated policy did not compile:\n%s\nerror: %v", text, err)
	}
}

func TestBuildCedarPolicy_PermitCompilesAndTargetsRoleAndActions(t *testing.T) {
	mp := &sentinelengine_models.ManagedPolicy{
		ID:     uuid.New(),
		Name:   "Property Manager",
		Effect: sentinelengine_models.EffectPermit,
		// intentionally unsorted to prove deterministic ordering
		Actions: []string{string(sentinel_actions.CreateProperty), string(sentinel_actions.ViewProperty)},
	}

	text, ok := buildCedarPolicy("ceo", mp)
	if !ok {
		t.Fatal("expected ok=true for a policy with actions")
	}
	mustParse(t, text)

	if !strings.Contains(text, `principal in Role::"ceo"`) {
		t.Errorf("policy must target the role; got:\n%s", text)
	}
	if !strings.Contains(text, `Action::"createProperty"`) || !strings.Contains(text, `Action::"viewProperty"`) {
		t.Errorf("policy must list both actions; got:\n%s", text)
	}
	if !strings.HasPrefix(text, "permit (") {
		t.Errorf("permit policy must start with permit; got:\n%s", text)
	}
	// Sorted: createProperty before viewProperty.
	if strings.Index(text, "createProperty") > strings.Index(text, "viewProperty") {
		t.Errorf("actions must be sorted for stable output; got:\n%s", text)
	}
	// Tenant guard present.
	if !strings.Contains(text, "resource.orgId == principal.orgId") {
		t.Errorf("policy must carry the tenant guard; got:\n%s", text)
	}
}

func TestBuildCedarPolicy_ForbidCompiles(t *testing.T) {
	mp := &sentinelengine_models.ManagedPolicy{
		Name:    "Block Delete",
		Effect:  sentinelengine_models.EffectForbid,
		Actions: []string{string(sentinel_actions.DeleteDocument)},
	}
	text, ok := buildCedarPolicy("intern", mp)
	if !ok {
		t.Fatal("expected ok=true")
	}
	mustParse(t, text)
	if !strings.HasPrefix(text, "forbid (") {
		t.Errorf("forbid effect must produce a forbid policy; got:\n%s", text)
	}
}

func TestBuildCedarPolicy_NoActionsIsSkipped(t *testing.T) {
	mp := &sentinelengine_models.ManagedPolicy{Name: "Empty", Effect: sentinelengine_models.EffectPermit}
	if _, ok := buildCedarPolicy("ceo", mp); ok {
		t.Error("a managed policy with no actions must produce nothing (ok=false)")
	}
}
