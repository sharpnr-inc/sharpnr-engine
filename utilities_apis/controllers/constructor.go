package utilitiesApi_controllers

import utilitiesApi_service "sharpnr-engine/utilities_apis/service"

type utilitiesApiController struct {
	utilitiesApiService utilitiesApi_service.UtilitiesAPIServiceInterface
}

func NewUtilitiesApiController(
	utilitiesApiService utilitiesApi_service.UtilitiesAPIServiceInterface,
) *utilitiesApiController {
	return &utilitiesApiController{
		utilitiesApiService: utilitiesApiService,
	}
}
