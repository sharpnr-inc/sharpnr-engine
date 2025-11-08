package sharphrcontrollers

import sharphrinterface "sharpnr-engine/modules/hr/sharphr_interface"

type sharpHrControllers struct {
	service sharphrinterface.SharpHrEmployeeService
}

func NewSharpHrEmployeeController(service sharphrinterface.SharpHrEmployeeService) *sharpHrControllers {
	return &sharpHrControllers{
		service: service,
	}
}
