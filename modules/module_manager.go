package modules

import (
	"fmt"
	"sharpnr-engine/core/constants"
	"sharpnr-engine/core/models"
)

type IModule interface {
	Initialize(org *models.Organization) error

	Type() constants.OrgType
}

type ModuleManager struct {
	modules    map[constants.OrgType]IModule
	baseModule IModule
}

func NewModuleManager() *ModuleManager {
	return &ModuleManager{
		modules: make(map[constants.OrgType]IModule),
	}
}

func (m *ModuleManager) InitializeModulesForOrg(org *models.Organization) error {
	// 1. Always initialize the base module first.
	if m.baseModule == nil {
		return fmt.Errorf("critical: base module is not registered")
	}
	fmt.Printf("Initializing BASE module for organization: %s\n", org.Name)
	if err := m.baseModule.Initialize(org); err != nil {
		return fmt.Errorf("failed to initialize base module for org %s: %w", org.SharpnrID, err)
	}
	return nil
}
