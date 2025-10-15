package policyengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type organisationPolicyRepository struct {
	db *gorm.DB
}

func NewOrgPolicyRepo(db *gorm.DB) *organisationPolicyRepository {
	return &organisationPolicyRepository{
		db: database.DB,
	}
}

// CreatePolicyOfOrg implements policyengine.PolicyRepository.
func (o *organisationPolicyRepository) CreatePolicyOfOrg(ctx context.Context, orgId, sharpnr_id, employeeId string) error {
	// ✅ Just delegate to the utils
	return CreateOrgDefaultPolicy(ctx, o.db, orgId, sharpnr_id, employeeId)
}

// GetOrgPolicyByPolicyId implements policyengine.PolicyRepository.
func (o *organisationPolicyRepository) GetOrgPolicyByPolicyId() {
	panic("unimplemented")
}

// GetEmployeePolicyByEmployeeId fetches the policy JSON and unmarshals it
func (o *organisationPolicyRepository) GetEmployeePolicyByEmployeeId(employeeID string) (*EmployeePolicy, error) {
	var orgPolicy models.OrganisationPolicy

	query := `
        (
            policy->'ORG_POLICY'->'EMPLOYEE_ACCESS' @> ?::jsonb
            OR policy->'ORG_POLICY'->'ORG_OWNERS' ? (?::text)
            OR policy->'ORG_POLICY'->'SUPER_ADMINS' ? (?::text)
        )
    `
	jsonCondition := fmt.Sprintf(`[{"EmployeeID": "%s"}]`, employeeID)

	err := o.db.
		Model(&models.OrganisationPolicy{}).
		Where(query, jsonCondition, employeeID, employeeID).
		Order("id").
		Limit(1).
		First(&orgPolicy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching policy: %w", err)
	}

	// Unmarshal the full JSON into OrgDefaultPolicy (your defined struct)
	var orgDefault OrgDefaultPolicy
	if err := json.Unmarshal(orgPolicy.Policy, &orgDefault); err != nil {
		return nil, fmt.Errorf("error unmarshalling policy json: %w", err)
	}

	// Extract the employee’s policy
	var empPolicy *EmployeePolicy
	for _, emp := range orgDefault.Policy.EMPLOYEE_ACCESS {
		if emp.EmployeeID == employeeID {
			empPolicy = &EmployeePolicy{
				EmployeeID: emp.EmployeeID,
				Privilages: map[string]interface{}{
					"SharpAccounting": emp.Privileges.SharpAccounting,
					"SharpHR":         emp.Privileges.SharpHR,
					"SharpHotel":      emp.Privileges.SharpHotel,
				},
			}
			break
		}
	}

	// If not in EMPLOYEE_ACCESS, check if SuperAdmin or OrgOwner
	if empPolicy == nil {
		role := ""
		for _, id := range orgDefault.Policy.SuperAdmins {
			if id == employeeID {
				role = "SUPER_ADMIN"
				break
			}
		}
		if role == "" {
			for _, id := range orgDefault.Policy.OrgOwners {
				if id == employeeID {
					role = "ORG_OWNER"
					break
				}
			}
		}

		if role != "" {
			empPolicy = &EmployeePolicy{
				EmployeeID: employeeID,
				Privilages: map[string]interface{}{
					"Role": role,
				},
			}
		}
	}

	if empPolicy == nil {
		return nil, nil // employee not found anywhere in policy
	}

	return empPolicy, nil
}

func (r *organisationPolicyRepository) GetEmployeePolicy(sharpnrID string, employeeID string) (*EmployeeAccess, error) {
	var policy models.OrganisationPolicy

	// 1️⃣ Query JSON metadata by SharpNRID
	if err := r.db.
		Where("policy -> 'Metadata' ->> 'SHARPNR_ID' = ?", sharpnrID).
		First(&policy).Error; err != nil {
		return nil, err
	}

	// 2️⃣ Parse JSON
	var parsed OrgDefaultPolicy
	if err := json.Unmarshal(policy.Policy, &parsed); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	// 3️⃣ Loop through EmployeeAccess to find the requested employee
	for _, emp := range parsed.Policy.EMPLOYEE_ACCESS {
		if emp.EmployeeID == employeeID {
			return &emp, nil
		}
	}

	return nil, errors.New("employee not found in org policy")
}
