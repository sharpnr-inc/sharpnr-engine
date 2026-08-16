package hrmsEngineControllers

import hrmsEngineService "sharpnr-engine/hrms_engine/service"

type hrmsControllers struct {
	hrmsService hrmsEngineService.HrmsEngineServiceInterface
}

func NewHrmsControllers(hrmsService hrmsEngineService.HrmsEngineServiceInterface) *hrmsControllers {
	return &hrmsControllers{
		hrmsService: hrmsService,
	}
}
