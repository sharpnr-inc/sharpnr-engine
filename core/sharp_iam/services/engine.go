package sharpIam_services

import (
	"context"
	"fmt"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// Namespace helps keep our Cedar types organized (e.g., Org::Employee)
const Namespace = "Org"

type SharpIAM struct {
	basePolicy []byte
}

func NewSharpIAM(basePolicy []byte) *SharpIAM {
	return &SharpIAM{
		basePolicy: basePolicy,
	}
}

func (e *SharpIAM) IsAuthorized(
	ctx context.Context,
	employeeID uuid.UUID,
	roleNames []string,
	action string,
	resType string,
	resID uuid.UUID,
	parentIDs []uuid.UUID,
) (cedar.Decision, error) {

	// 1️⃣ Parse Policy
	var policy cedar.Policy
	if err := policy.UnmarshalCedar(e.basePolicy); err != nil {
		return false, fmt.Errorf("bad policy syntax: %w", err)
	}

	ps := cedar.NewPolicySet()
	ps.Add("policy0", &policy)

	// 2️⃣ Build Entities
	entities := cedar.EntityMap{}

	cedar.NewPolicySet()

	// ---------- Principal ----------
	pUID := cedar.NewEntityUID(
		cedar.EntityType(fmt.Sprintf("%s::Employee", Namespace)),
		cedar.String(employeeID.String()),
	)

	var roleUIDs []cedar.EntityUID
	for _, r := range roleNames {
		roleUIDs = append(roleUIDs,
			cedar.NewEntityUID(
				cedar.EntityType(fmt.Sprintf("%s::Role", Namespace)),
				cedar.String(r),
			),
		)
	}

	entities[pUID] = cedar.Entity{
		UID:     pUID,
		Parents: cedar.NewEntityUIDSet(roleUIDs...), // ✅ FIX HERE
	}

	// ---------- Resource ----------
	rUID := cedar.NewEntityUID(
		cedar.EntityType(fmt.Sprintf("%s::%s", Namespace, resType)),
		cedar.String(resID.String()),
	)

	var parentUIDs []cedar.EntityUID
	for _, p := range parentIDs {
		parentUIDs = append(parentUIDs,
			cedar.NewEntityUID(
				cedar.EntityType(fmt.Sprintf("%s::Organization", Namespace)),
				cedar.String(p.String()),
			),
		)
	}

	entities[rUID] = cedar.Entity{
		UID:     rUID,
		Parents: cedar.NewEntityUIDSet(parentUIDs...), // ✅ FIX HERE
	}

	// 3️⃣ Create Request
	req := cedar.Request{
		Principal: pUID,
		Action: cedar.NewEntityUID(
			cedar.EntityType(Namespace+"::Action"),
			cedar.String(action),
		),
		Resource: rUID,
		Context:  cedar.NewRecord(cedar.RecordMap{}),
	}

	// 4️⃣ Authorize
	decision, diagnostics := cedar.Authorize(ps, entities, req)

	if len(diagnostics.Errors) > 0 {
		return false, fmt.Errorf("authorization failed: %v", diagnostics.Errors)
	}

	return decision, nil
}
