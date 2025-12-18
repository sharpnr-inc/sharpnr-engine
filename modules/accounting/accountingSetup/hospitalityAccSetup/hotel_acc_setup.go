package hospitality_acc_setup

import "context"

type HotelAccountingSetupStep struct{}

func (s *HotelAccountingSetupStep) Name() string {
	return "BarAccountingSetup"
}

func (s *HotelAccountingSetupStep) Execute(
	ctx context.Context,
	c AccSetupContext,
) error {
	panic("unimplemented")
}
