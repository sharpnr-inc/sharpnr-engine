package sharpnr_policy_engine

import "github.com/google/uuid"

// MVP / stub implementation
type DefaultPolicyResolver struct{}

func NewDefaultPolicyResolver() *DefaultPolicyResolver {
	return &DefaultPolicyResolver{}
}

func (r *DefaultPolicyResolver) ResolveEmployeePolicy(
	employeeID uuid.UUID,
) (ResolvedPolicy, error) {

	// This will later come from DB
	return ResolvedPolicy{
		AllowedModules: map[string]bool{
			"bar":        true,
			"restaurant": false,
			"hotel":      false,
		},
		AllowedActions: map[string]map[string]bool{
			"bar": {
				"menu":   true,
				"orders": true,
			},
		},
	}, nil
}
