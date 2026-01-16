package sharpnr_policy_engine

import "github.com/google/uuid"

// Interface – real implementation can be DB-backed
type PolicyResolver interface {
	ResolveEmployeePolicy(employeeID uuid.UUID) (ResolvedPolicy, error)
}
