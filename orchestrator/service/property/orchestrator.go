package property_orchestrator

import (
	org_property_services "sharpnr-engine/core/org_property/services"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"
	propertyHospitalityOrchestrator "sharpnr-engine/orchestrator/service/property/hospitality"
	dbtx "sharpnr-engine/pkg/dbTx"
)

// PropertyOrchestrator implements
// orchestrator_service.OrganizationPropertyOrchestrationInterface: property CRUD,
// plus attaching and listing the business/infra modules of a property.
//
// Property flows are guarded by the Sentinel engine. Action and resource-type
// ids come from the org-property module's mod_resources so checks and the
// policy-builder palette can never drift apart.
type PropertyOrchestrator struct {
	guard               *orchestrator_authz.Guard
	txManager           dbtx.TxManager
	orgPropertyService  org_property_services.OrganizationPropertyServicesInterfaces
	organizationService organization_services.OrganizationServiceInterface

	// One embedded orchestrator per business domain. Their methods are promoted
	// onto PropertyOrchestrator and on up to the composed facade.
	*propertyHospitalityOrchestrator.HospitalityOrchestrator
	// soon: *healthcare_orchestrator.HealthcareOrchestrator
	// soon: *retail_orchestrator.RetailOrchestrator
}

func New(
	guard *orchestrator_authz.Guard,
	txManager dbtx.TxManager,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	organizationService organization_services.OrganizationServiceInterface,
	hospitality *propertyHospitalityOrchestrator.HospitalityOrchestrator,
) *PropertyOrchestrator {
	return &PropertyOrchestrator{
		guard:                   guard,
		txManager:               txManager,
		orgPropertyService:      orgPropertyService,
		organizationService:     organizationService,
		HospitalityOrchestrator: hospitality,
	}
}
