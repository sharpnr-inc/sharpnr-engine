package stockwaveEngineModels

import (
	"time"

	"github.com/google/uuid"
)

type OrgWarehouse struct {
	ID                           uuid.UUID
	Name                         string
	OrganizationID               uuid.UUID
	OrgStockwaveEngineAccountsID uuid.UUID
	Description                  string
	Location                     string
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
	CostCenterID                 *uuid.UUID
}

func (OrgWarehouse) TableName() string {
	return "stockwave_engine.org_warehouse"
}

type OrgWarehouseSettings struct {
	ID             uuid.UUID
	OrgWarehouseID uuid.UUID
}
