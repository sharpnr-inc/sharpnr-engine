package app

import (
	utlitiesApi_repo "sharpnr-engine/utilities_apis/repo"
	utilitiesApi_service "sharpnr-engine/utilities_apis/service"

	"gorm.io/gorm"
)

func bootstrapUtilitiesApiService(db *gorm.DB) utilitiesApi_service.UtilitiesAPIServiceInterface {
	utilitiesApiRepo := utlitiesApi_repo.NewUtilitiesApiRepo(db)
	return utilitiesApi_service.NewUtilitesApiService(
		utilitiesApiRepo,
	)

}
