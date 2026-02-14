package orchestrator_service

import(
	"context"
	"github.com/google/uuid"
	"sharpnr-engine/app/hospitality_domain"
	"sharpnr-engine/core/dbTx"
	"sharpnr-engine/core/org_legal_entity/service"
	"sharpnr-engine/core/org_property/services"
	"sharpnr-engine/core/org_setup_wizard/services"
	"sharpnr-engine/core/organization/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/service_engines/accounting_engine/services"
	"testing"
)

func Test_orchestratorService_GetPropertiesAndRestaurant(t *testing.T) {
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
		orgID   uuid.UUID
		empID   uuid.UUID
		want    *orchestrator_dtos.PropertyOfResturants
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSharpnrOrchestratorService(tt.organizationService, tt.orgPropertyService, tt.orgSetupWizService, tt.hospitalityDomainSvc, tt.accountingEngineService, tt.legalEntityService, tt.txManager)
			got, gotErr := s.GetPropertiesAndRestaurant(context.Background(), tt.orgID, tt.empID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetPropertiesAndRestaurant() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetPropertiesAndRestaurant() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetPropertiesAndRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
