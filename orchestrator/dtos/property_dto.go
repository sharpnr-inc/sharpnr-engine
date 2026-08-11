package orchestrator_dtos

import (
	general_infra_utils "sharpnr-engine/property_business_ops/general_infra/utils"
	healthcare_utilities "sharpnr-engine/property_business_ops/healthcare/utility"
	hospitality_utilities "sharpnr-engine/property_business_ops/hospitality/utilities"

	"github.com/google/uuid"
)

// --------------------------------------------------------------------------------------//

type PropertiesOfOrganization_ResDto struct {
	Properties []PropertiesReaderSlice `json:"properties"`
}

type PropertiesReaderSlice struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AddressLine string    `json:"address_line"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Code        string    `json:"code"`
}

// --------------------------------------------------------------------------------------//

type PropertyDetails_res_dto struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	City        string `json:"city"`
	AddressLine string `json:"address_line"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

// --------------------------------------------------------------------------------------//

type CreatePropertyReq_dto struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
}

/* ------------------------------------ . ----------------------------------- */

type HospitalityModuleQuantity struct {
	Module   hospitality_utilities.HospitalityModulesType `json:"module" binding:"required"`
	Quantity int                                          `json:"quantity" binding:"required,min=1"`
}

type HealthCareModuleQuantity struct {
	Module   healthcare_utilities.HealthcareModulesType `json:"module" binding:"required"`
	Quantity int                                        `json:"quantity" binding:"required,min=1"`
}

// AddBusinessOpModulesRequestDto is the request body for attaching business
// operational (revenue-generating) modules to a property.
//
// Valid "module" values:
//   - hospitality_modules: RESTAURANT, HOTEL, BAR, NIGHT_CLUB, RESORT, CLOUD_KITCHEN/TAKEAWAYS
//   - healthcare_modules:  HOSPITAL, CLINIC, PHARMACY, DIAGNOSTICS, TELEMEDICINE, HOME_CARE
//
// Example request:
//
//	{
//	  "property_id": "3f0c9b1e-4a2d-4c8e-9f7a-1b2c3d4e5f60",
//	  "hospitality_modules": [
//	    {"module": "RESTAURANT", "quantity": 2},
//	    {"module": "BAR", "quantity": 1}
//	  ],
//	  "healthcare_modules": [
//	    {"module": "PHARMACY", "quantity": 1}
//	  ]
//	}
type AddBusinessOpModulesRequestDto struct {
	PropertyID         uuid.UUID                   `json:"property_id"`
	HospitalityModules []HospitalityModuleQuantity `json:"hospitality_modules"`
	HealthcareModules  []HealthCareModuleQuantity  `json:"healthcare_modules"`
}

/* ------------------------------------ . ----------------------------------- */
type HospitalityInfraModuleQuantity struct {
	Module   hospitality_utilities.HospitalityInfraModulesTypes `json:"module" binding:"required"`
	Quantity int                                                `json:"quantity" binding:"required,min=1"`
}

type HealthCareInfraModuleQuantity struct {
	Module   healthcare_utilities.HealthcareInfraModulesType `json:"module" binding:"required"`
	Quantity int                                             `json:"quantity" binding:"required,min=1"`
}

type GeneralInfraModuleQuantity struct {
	Module   general_infra_utils.GeneralInfraModulesType `json:"module" binding:"required"`
	Quantity int                                         `json:"quantity" binding:"required,min=1"`
}

// AddInfraModulesRequestDto is the request body for attaching infrastructure
// (physical / operational area) modules to a property.
//
// Valid "module" values:
//   - general_modules:   "Front Desk / Reception", "Billing Counter", "Security Desk"
//   - hospitality_infra: KITCHEN, BAR_COUNTER, DINING_HALL, INVENTORY_STORE, RECEPTION_FRONT_DESK, HOUSEKEEPING
//   - health_care_infra: WARD_INPATIENT_UNIT, OPERATION_THEATRE, LABORATORY, PHARMACY_COUNTER, RECEPTION_REGISTRATION, RADIOLOGY_IMAGING
//
// Example request:
//
//	{
//	  "property_id": "3f0c9b1e-4a2d-4c8e-9f7a-1b2c3d4e5f60",
//	  "general_modules": [
//	    {"module": "Front Desk / Reception", "quantity": 1},
//	    {"module": "Billing Counter", "quantity": 1}
//	  ],
//	  "hospitality_infra": [
//	    {"module": "KITCHEN", "quantity": 2},
//	    {"module": "BAR_COUNTER", "quantity": 1}
//	  ],
//	  "health_care_infra": [
//	    {"module": "WARD_INPATIENT_UNIT", "quantity": 3},
//	    {"module": "OPERATION_THEATRE", "quantity": 1}
//	  ]
//	}
type AddInfraModulesRequestDto struct {
	PropertyID       uuid.UUID                        `json:"property_id"`
	GeneralModules   []GeneralInfraModuleQuantity     `json:"general_modules"`
	HospitalityInfra []HospitalityInfraModuleQuantity `json:"hospitality_infra"`
	HealthCareInfra  []HealthCareInfraModuleQuantity  `json:"health_care_infra"`
}

/* ------------------------------------ . ----------------------------------- */

type HospitalityBusinessModule struct {
	ID           uuid.UUID                                    `json:"id"`
	Name         string                                       `json:"name"`
	ModuleType   hospitality_utilities.HospitalityModulesType `json:"module_type"`
	Status       bool                                         `json:"status"`
	IsConfigured bool                                         `json:"is_configured"`
}

type HealthCareBusinessModule struct {
	ID           uuid.UUID                                  `json:"id"`
	Name         string                                     `json:"name"`
	ModuleType   healthcare_utilities.HealthcareModulesType `json:"module_type"` // was hospitality type
	Status       bool                                       `json:"status"`
	IsConfigured bool                                       `json:"is_configured"`
}

// Response body for all the Business Operation Modules of a Property.
type PropertyBusinessOperationModulesResponseDto struct {
	HospitalityModules []HospitalityBusinessModule `json:"hospitality_business_operation_modules"`
	HealthCareModules  []HealthCareBusinessModule  `json:"healthcare_business_operation_modules"`
}

type GeneralInfraModule struct {
	ID        uuid.UUID                                   `json:"id"`
	Name      string                                      `json:"name"`
	InfraType general_infra_utils.GeneralInfraModulesType `json:"infra_type"`
}

type HospitalityInfraModule struct {
	ID        uuid.UUID                                          `json:"id"`
	Name      string                                             `json:"name"`
	Status    bool                                               `json:"status"`
	InfraType hospitality_utilities.HospitalityInfraModulesTypes `json:"infra_type"`
	Meta      interface{}                                        `json:"meta,omitempty"` // shape TBD
}

type HealthCareInfraModule struct {
	ID        uuid.UUID                                       `json:"id"`
	Name      string                                          `json:"name"`
	InfraType healthcare_utilities.HealthcareInfraModulesType `json:"infra_type"`
}

type PropertyInfraModulesResponseDto struct {
	GeneralInfraModules     []GeneralInfraModule     `json:"general_infra_modules"`
	HospitalityInfraModules []HospitalityInfraModule `json:"hospitality_infra_modules"`
	HealthCareInfraModules  []HealthCareInfraModule  `json:"healthcare_infra_modules"`
}
