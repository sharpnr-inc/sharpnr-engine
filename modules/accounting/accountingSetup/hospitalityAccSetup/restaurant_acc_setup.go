package hospitality_acc_setup

import "context"

type RestaurantAccountingSetupStep struct{}

func (s *RestaurantAccountingSetupStep) Name() string {
	return "BarAccountingSetup"
}

func (s *RestaurantAccountingSetupStep) Execute(
	ctx context.Context,
	c AccSetupContext,
) error {
	panic("unimplemented")
}
