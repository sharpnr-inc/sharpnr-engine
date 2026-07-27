package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"
	"time"

	"github.com/google/uuid"
)

func (d *employeeRepo) CreateEmployeeProvision(ctx context.Context, orgID uuid.UUID, provisions []coreEmployeeModels.EmployeeProvisioning) error {
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (d *employeeRepo) GetEmployeeProvioionByEmployeeID(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) ([]coreEmployeeModels.EmployeeProvisioning, error)

/* ---------------------------------- <***> --------------------------------- */

func (d *employeeRepo) AddConsumeTimeInProvisioningToken(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID, consumptionTime time.Time) error
