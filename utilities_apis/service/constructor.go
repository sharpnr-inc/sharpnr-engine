package utilitiesApi_service

import utlitiesApi_repo "sharpnr-engine/utilities_apis/repo"

type utilitiesApiService struct {
	repo utlitiesApi_repo.UtilitesAPIRepositoryInterface
}

func NewUtilitesApiService(
	repo utlitiesApi_repo.UtilitesAPIRepositoryInterface,
) *utilitiesApiService {
	return &utilitiesApiService{
		repo: repo,
	}
}
