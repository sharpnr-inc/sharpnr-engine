package modules

import (
	"fmt"

	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

// IModule defines the standard interface for all business modules.
// Each module (base, hotel, retail, etc.) must implement this interface
// to be managed by the ModuleManager.
type IModule interface {
	// Initialize sets up the module-specific data for a new organization.
	// This includes creating default departments, roles, or other necessary records.
	Initialize(org *models.Organization) error

	// Type returns the specific OrgType the module is responsible for.
	// For the base module, this can be a special value like constants.OrgTypeBase.
	Type() constants.OrgType
}

// ModuleManager is the central coordinator responsible for managing the lifecycle
// of all business modules in the application.
type ModuleManager struct {
	// modules is a map where keys are constants.OrgType and values are the
	// corresponding module implementations.
	modules map[constants.OrgType]IModule
	// baseModule is the universal module that is initialized for every organization.
	baseModule IModule
}

// NewModuleManager creates and returns a new ModuleManager instance.
// It initializes the internal map used to store and look up modules.
func NewModuleManager() *ModuleManager {
	return &ModuleManager{
		modules: make(map[constants.OrgType]IModule),
	}
}

// RegisterModule allows a module to be registered with the manager.
// This should be called at application startup for each available module.
// If a module is the base module, it is stored separately.
func (m *ModuleManager) RegisterModule(module IModule) {
	if module.Type() == constants.OrgTypeBase {
		m.baseModule = module
		fmt.Printf("Registered BASE module.\n")
	} else {
		m.modules[module.Type()] = module
		fmt.Printf("Registered module for type: %s\n", module.Type())
	}
}

// InitializeModulesForOrg is the core function that sets up a new organization.
// It is called by the OrganizationService after a new organization is created.
// It ensures that the universal base module is initialized first, followed by the
// specific module that matches the organization's operation type.
func (m *ModuleManager) InitializeModulesForOrg(org *models.Organization) error {
	// 1. Always initialize the base module first.
	if m.baseModule == nil {
		return fmt.Errorf("critical: base module is not registered")
	}
	fmt.Printf("Initializing BASE module for organization: %s\n", org.Name)
	if err := m.baseModule.Initialize(org); err != nil {
		return fmt.Errorf("failed to initialize base module for org %s: %w", org.SharpnrOrgID, err)
	}

	// 2. Find and initialize the specific module for the organization's type.
	specificModule, exists := m.modules[org.OrgOprationType]
	if !exists {
		// This is not necessarily an error. Some org types might only have a base module.
		fmt.Printf("No specific module found for type '%s'. Only base module was initialized.\n", org.OrgOprationType)
		return nil
	}

	fmt.Printf("Initializing module '%s' for organization: %s\n", specificModule.Type(), org.Name)
	if err := specificModule.Initialize(org); err != nil {
		return fmt.Errorf("failed to initialize module %s for org %s: %w", specificModule.Type(), org.SharpnrOrgID, err)
	}

	return nil
}
