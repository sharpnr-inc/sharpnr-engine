package orchestrator_service

import (
	"context"
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	dbtx "sharpnr-engine/core/dbTx"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	"testing"

	"github.com/google/uuid"
)

func Test_orchestratorService_ConfigureHotel(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		organizationService     organization_services.OrganizationServiceInterface
		orgPropertyService      org_property_services.OrganizationPropertyServicesInterfaces
		orgSetupWizService      org_setup_wizard_services.OrgSetupWizardServiceInterface
		hospitalityDomainSvc    *hospitalitydomain.HospitalityDomain
		accountingEngineService accounting_service.AccountingServiceInterface
		legalEntityService      org_legal_entity_service.OrgLegalEntityServiceInterface
		txManager               dbtx.TxManager
		// Named input parameters for target function.
		orgID      uuid.UUID
		propertyID uuid.UUID
		req        *orchestrator_dtos.ConfigureHotel_Request_dto
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSharpnrOrchestratorService(tt.organizationService, tt.orgPropertyService, tt.orgSetupWizService, tt.hospitalityDomainSvc, tt.accountingEngineService, tt.legalEntityService, tt.txManager)
			gotErr := c.ConfigureHotel(context.Background(), tt.orgID, tt.propertyID, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ConfigureHotel() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ConfigureHotel() succeeded unexpectedly")
			}
		})
	}
}

func Test_orchestratorService_GetPropertyInfra(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		organizationService     organization_services.OrganizationServiceInterface
		orgPropertyService      org_property_services.OrganizationPropertyServicesInterfaces
		orgSetupWizService      org_setup_wizard_services.OrgSetupWizardServiceInterface
		hospitalityDomainSvc    *hospitalitydomain.HospitalityDomain
		accountingEngineService accounting_service.AccountingServiceInterface
		legalEntityService      org_legal_entity_service.OrgLegalEntityServiceInterface
		txManager               dbtx.TxManager
		// Named input parameters for target function.
		propertyID uuid.UUID
		orgID      uuid.UUID
		infraType  *string
		want       *orchestrator_dtos.ListAllInfraOfProperty_Res_dto
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSharpnrOrchestratorService(tt.organizationService, tt.orgPropertyService, tt.orgSetupWizService, tt.hospitalityDomainSvc, tt.accountingEngineService, tt.legalEntityService, tt.txManager)
			got, gotErr := s.GetPropertyInfra(context.Background(), tt.propertyID, tt.orgID, tt.infraType)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetPropertyInfra() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetPropertyInfra() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetPropertyInfra() = %v, want %v", got, tt.want)
			}
		})
	}
}
