package policyengine

type policySetup struct {
	repo PolicyRepository
}

func NewPolicySetup(repo PolicyRepository) PolicySetup {
	return &policySetup{repo: repo}
}

func (ps *policySetup) CreateOrgDefaultPolicy(orgID, orgType string) (*OrgDefaultPolicy, error) {
	return nil, nil
}

func (ps *policySetup) UpdatePolicy(orgID string) (*OrgDefaultPolicy, error) {
	return nil, nil
}
