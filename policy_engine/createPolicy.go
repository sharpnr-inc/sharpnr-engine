package policyengine

import (
	"context"
	"sharpnr-engine/core/models"

	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrgDefaultPolicy(ctx context.Context, db *gorm.DB, orgID, sharpnr_id, employeeID string) error {
	// 1️⃣ Create default policy structure
	defaultPrivileges := EmployeeAccess{
		EmployeeID: employeeID,
		Privileges: EmployeePrivileges{
			SharpAccounting: &ModulePrivileges{
				Access:        "NONE",
				Transactions:  []string{},
				Invoice:       []string{},
				BalanceSheets: []string{},
			},
			SharpHR: &ModulePrivileges{
				Access: "NONE",
			},
			SharpHotel: &ModulePrivileges{
				Access: "NONE",
			},
		},
	}

	orgPolicy := OrgDefaultPolicy{
		Metadata: Metadata{
			OrgID:     orgID,
			SharpnrID: sharpnr_id,
			PolicyID:  uuid.NewString(),
			Version:   1.0,
			OrgType:   "DEFAULT",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		Policy: Policy{
			SuperAdmins:     []string{employeeID},
			OrgOwners:       []string{employeeID},
			EMPLOYEE_ACCESS: []EmployeeAccess{defaultPrivileges},
		},
	}

	// 2️⃣ Marshal to JSON
	policyJSON, err := json.Marshal(orgPolicy)
	if err != nil {
		return err
	}

	orgUUID, err := uuid.Parse(orgID)
	if err != nil {
		return err
	}

	// 3️⃣ Save in DB
	model := models.OrganizationPolicy{
		ID:             uuid.New(),
		OrganizationID: orgUUID,
		Policy:         policyJSON,
		CreatedAt:      time.Now().Format(time.RFC3339),
		UpdatedAt:      time.Now().Format(time.RFC3339),
	}

	return db.WithContext(ctx).Create(&model).Error
}
