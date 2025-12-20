package hospitality_setup

import "context"

type HospitalitySetupStepInterface interface {
	Name() string
	Execute(ctx context.Context, c SetupContext) error
}
