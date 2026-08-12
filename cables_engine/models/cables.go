// this cables main models
package cables_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// A organization only gonna have one Cables Account
// the reason for this Account tables because one day we might need to Individual Cables Account

type OrgCablesAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	IsActive       bool // this is switch which will or can turn-off the Cables service in the organization
}

func (t *OrgCablesAccount) TableName() string {
	return "cables.org_cables_accounts"
}

/* ---------------------------------- <***> --------------------------------- */
