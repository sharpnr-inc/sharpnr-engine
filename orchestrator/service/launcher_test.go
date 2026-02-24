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
	"sharpnr-engine/sero"
	"sharpnr-engine/service_engines/accounting_engine/services"
	"testing"
)

func Test_orchestratorService_GetSharpnrLauncherData(t *testing.T) {
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
		seroEngine              *sero.SEROEngine
		// Named input parameters for target function.
		orgID   uuid.UUID
		emplID  uuid.UUID
		want    orchestrator_dtos.OrgLauncherResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSharpnrOrchestratorService(tt.organizationService, tt.orgPropertyService, tt.orgSetupWizService, tt.hospitalityDomainSvc, tt.accountingEngineService, tt.legalEntityService, tt.txManager, tt.seroEngine)
			got, gotErr := c.GetSharpnrLauncherData(context.Background(), tt.orgID, tt.emplID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetSharpnrLauncherData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetSharpnrLauncherData() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetSharpnrLauncherData() = %v, want %v", got, tt.want)
			}
		})
	}
}
