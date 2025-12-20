package hospitality_acc_setup

import "context"

type HospitalityAccountingSetupStepInterface interface {
	Name() string
	Execute(ctx context.Context, c AccSetupContext) error
}
