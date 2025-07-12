package base

import (
	"fmt"

	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

// baseModule implements the ModuleInitializer interface for base organization setup.
type baseModule struct {
	departmentService interfaces.DepartmentService // Assuming this interface exists
	terminalService   interfaces.TerminalService   // Assuming this interface exists
}

// NewBaseModule creates a new instance of the base module initializer.
// It requires services for department and terminal management.
func NewBaseModule(deptService interfaces.DepartmentService, termService interfaces.TerminalService) ModuleInitializer {
	return &baseModule{
		departmentService: deptService,
		terminalService:   termService,
	}
}

// InitializeOrganization performs base setup for a newly registered organization.
// This includes creating default departments and their associated terminals.
func (bm *baseModule) InitializeOrganization(org *models.Organization) error {
	fmt.Printf("Initializing Base module for organization: %s (UUID: %s)\n", org.Name, org.SharpnrOrgUUID)

	// TODO: Implement actual base initialization logic here.
	// 1. Create 'Accounting' Department
	// Example (conceptual):
	// accountingDept, err := bm.departmentService.CreateDefaultAccountingDepartment(org.SharpnrOrgUUID)
	// if err != nil {
	// 	return fmt.Errorf("failed to create accounting department: %w", err)
	// }

	// 2. Initialize default terminals for 'Accounting' Department
	// Example (conceptual):
	// err = bm.terminalService.InitializeDefaultAccountingTerminals(org.SharpnrOrgUUID, accountingDept.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to initialize accounting terminals: %w", err)
	// }

	fmt.Println("Base module initialization complete.")
	return nil
}
