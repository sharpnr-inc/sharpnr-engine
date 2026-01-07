package hospitality_setup

func HospitalityModuleSetupFactory(modules []string) []HospitalitySetupStepInterface {
	steps := []HospitalitySetupStepInterface{}

	for _, m := range modules {
		switch m {
		case "BAR":
			steps = append(steps, &BarSetupStep{})
		case "HOTEL":
			steps = append(steps, &HotelSetupStep{})
		case "RESTAURANT":
			steps = append(steps, &RestaurantSetupStep{})
		}
	}
	return steps
}
