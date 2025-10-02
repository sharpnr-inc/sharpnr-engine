package policyengine

import (
	"encoding/json"
	"fmt"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func CreateDefaultPolicy(org *models.Organization) (datatypes.JSON, error) {
	policy := OrgDefaultPolicy{
		Metadata: Metadata{
			OrgID:     org.ID.String(),
			SharpnrID: org.SharpnrID,
			PolicyID:  uuid.New().String(),
			Version:   1.0,
			OrgType:   string(org.OrgOprationType),
		},
	}

	b, err := json.Marshal(policy)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal policy: %w", err)
	}

	return datatypes.JSON(b), nil
}
