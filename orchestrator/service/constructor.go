package orchestrator_service

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
)

type orchestratorService struct {
	organizationService     organization_services.OrganizationServiceInterface
	orgPropertyService      org_property_services.OrganizationPropertyServicesInterfaces
	orgSetupWizService      org_setup_wizard_services.OrgSetupWizardServiceInterface
	hospitalityDomainSvc    *hospitalitydomain.HospitalityDomain
	accountingEngineService accounting_service.AccountingServiceInterface
	legalEntityService      legal_entity_service.OrgLegalEntityServiceInterface
	txManager               dbtx.TxManager
	seroEngine              *sero.SEROEngine
}

func NewSharpnrOrchestratorService(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	accountingEngineService accounting_service.AccountingServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *orchestratorService {
	return &orchestratorService{
		organizationService:     organizationService,
		orgPropertyService:      orgPropertyService,
		orgSetupWizService:      orgSetupWizService,
		hospitalityDomainSvc:    hospitalityDomainSvc,
		accountingEngineService: accountingEngineService,
		legalEntityService:      legalEntityService,
		txManager:               txManager,
		seroEngine:              seroEngine,
	}
}
