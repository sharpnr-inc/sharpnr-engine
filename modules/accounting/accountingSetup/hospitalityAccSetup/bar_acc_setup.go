package hospitality_acc_setup

import "context"

type BarAccountingSetupStep struct{}

func (s *BarAccountingSetupStep) Name() string {
	return "BarAccountingSetup"
}

func (s *BarAccountingSetupStep) Execute(
	ctx context.Context,
	c AccSetupContext,
) error {
	panic("unimplemented")
}
