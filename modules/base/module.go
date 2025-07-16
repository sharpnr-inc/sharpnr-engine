package base

// import (
// 	"fmt"

// 	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
// 	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
// 	"github.com/sharpnr-inc/sharpnr-engine/core/models"
// 	"github.com/sharpnr-inc/sharpnr-engine/modules"
// )

// // Ensure BaseModule implements the IModule interface at compile time.
// var _ modules.IModule = (*Module)(nil)

// // Module implements the IModule interface for universal, base-level features.
// // It is responsible for setting up components that every organization needs.
// type Module struct {
// 	// depRepo is the dependency used to interact with the department data store.
// 	depRepo interfaces.DepartmentRepository
// }

// // NewModule is a constructor that creates a new instance of the BaseModule.
// // It requires a DepartmentRepository to be injected for its operations.
// func NewModule(depRepo interfaces.DepartmentRepository) *Module {
// 	return &Module{depRepo: depRepo}
// }

// // Type returns the specific operation type for this module.
// // For the base module, it returns a special constant recognized by the ModuleManager.
// func (m *Module) Type() constants.OrgType {
// 	return m.Type()
// }

// // Initialize sets up the default, universal departments for a new organization.
// // It is called by the ModuleManager during the organization registration process.
// func (m *Module) Initialize(org *models.Organization) error {
// 	fmt.Printf("BASE MODULE: Initializing default departments for %s\n", org.Name)

// 	// Define the departments that every organization receives by default.
// 	defaultDepts := []models.Department{
// 		{
// 			Name:           "Human Resources",
// 			OrganizationID: org.ID, // Link the department to the new organization.
// 		},
// 		{
// 			Name:           "Accounting",
// 			OrganizationID: org.ID, // Link the department to the new organization.
// 		},
// 	}

// 	// Use the injected repository to persist each default department.
// 	for _, dept := range defaultDepts {
// 		// It's often good practice to check if a department with this name already exists
// 		// for this org, but for a clean initialization, we'll assume it doesn't.
// 		if err := m.depRepo.Create(&dept); err != nil {
// 			// If any department creation fails, we must stop and return the error
// 			// to prevent a partially configured organization.
// 			return fmt.Errorf("base module: failed to create department '%s' for org '%s': %w", dept.Name, org.SharpnrOrgID, err)
// 		}
// 	}

// 	fmt.Printf("BASE MODULE: Successfully initialized default departments for %s\n", org.Name)
// 	return nil
// }
