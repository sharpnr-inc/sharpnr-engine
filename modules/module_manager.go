package modules

import (
	"errors"
	"fmt"

	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"github.com/sharpnr-inc/sharpnr-engine/modules/base"
)

// IModuleManager manages the lifecycle of all available modules and their initializers.
type IModuleManager interface {
	// RegisterBaseInitializer registers a universal initializer that runs for all organizations.
	RegisterBaseInitializer(initializer base.ModuleInitializer)

	// RegisterModuleInitializer registers a ModuleInitializer for a specific organization type.
	RegisterModuleInitializer(orgType constants.OrgType, initializer base.ModuleInitializer)

	// InitializeOrganization calls the appropriate ModuleInitializer for the given organization.
	InitializeOrganization(org *models.Organization) error
}

type moduleManager struct {
	baseInitializer base.ModuleInitializer
	initializers    map[constants.OrgType]base.ModuleInitializer
}

// NewModuleManager creates a new instance of IModuleManager.
func NewModuleManager() IModuleManager {
	return &moduleManager{
		initializers: make(map[constants.OrgType]base.ModuleInitializer),
	}
}

// RegisterBaseInitializer registers a universal initializer that runs for all organizations.
func (mm *moduleManager) RegisterBaseInitializer(initializer base.ModuleInitializer) {
	mm.baseInitializer = initializer
}

// RegisterModuleInitializer registers a ModuleInitializer for a specific organization type.
func (mm *moduleManager) RegisterModuleInitializer(orgType constants.OrgType, initializer base.ModuleInitializer) {
	mm.initializers[orgType] = initializer
}

// InitializeOrganization calls the appropriate ModuleInitializer for the given organization.
func (mm *moduleManager) InitializeOrganization(org *models.Organization) error {
	if org == nil {
		return errors.New("organization cannot be nil")
	}

	// First, run the base initializer for all organizations
	if mm.baseInitializer != nil {
		if err := mm.baseInitializer.InitializeOrganization(org); err != nil {
			return fmt.Errorf("base module initialization failed: %w", err)
		}
	}

	// Then, run the specific module initializer if one exists for the organization type
	initializer, ok := mm.initializers[org.OrgOprationType]
	if !ok {
		// No specific initializer found for this organization type. This might be okay
		// if some organization types don't require specific module setup beyond the base.
		// For now, we'll just log or return nil if no specific initializer is found.
		// If you want to enforce that every OrgType must have a specific initializer,
		// you would return an error here.
		return nil // Or return an error if specific initializer is mandatory
	}

	return initializer.InitializeOrganization(org)
}
