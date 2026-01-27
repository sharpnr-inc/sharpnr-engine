package accounting_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accounting_engine_contracts "sharpnr-engine/sharp_services/accounting_engine/contracts"
	accounting_repo "sharpnr-engine/sharp_services/accounting_engine/repository"
	"testing"
)

func Test_accountingServices_CreateDraftAccountingSetup(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		accRepo   accounting_repo.AccountingRepositoryInterface
		txManager dbtx.TxManager
		// Named input parameters for target function.
		input   *accounting_engine_contracts.DraftAccountingSetupInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewAccountingService(tt.accRepo, tt.txManager)
			gotErr := r.CreateDraftAccountingSetup(context.Background(), tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateDraftAccountingSetup() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateDraftAccountingSetup() succeeded unexpectedly")
			}
		})
	}
}
