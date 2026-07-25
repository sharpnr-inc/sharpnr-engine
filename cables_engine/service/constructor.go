package cables_service

import cables_repo "sharpnr-engine/cables_engine/repo"

type cablesService struct {
	cablesRepo cables_repo.CablesRepoInterface
	// and other services
}
