package iamEngineControllers

import iamEngineService "sharpnr-engine/iam_engine/service"

type iamEngineControllers struct {
	iamEngineService iamEngineService.IamEngineServiceInterface
}

func NewIamEngineServiceController(iamEngineService iamEngineService.IamEngineServiceInterface) *iamEngineControllers {
	return &iamEngineControllers{
		iamEngineService: iamEngineService,
	}
}
