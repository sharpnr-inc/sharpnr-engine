package hospitality_acc_setup

func HospitalityAccountingSetupFactory(
	subModules []string,
) []HospitalityAccountingSetupStepInterface {
	steps := []HospitalityAccountingSetupStepInterface{}

	for _, m := range subModules {
		switch m {
		case "BAR":
			steps = append(steps, &BarAccountingSetupStep{})
		case "HOTEL":
			steps = append(steps, &HotelAccountingSetupStep{})
		case "RESTAURANT":
			steps = append(steps, &RestaurantAccountingSetupStep{})
		}

	}
	return steps
}
