package app

import (
	"sharpnr-engine/core/sharp_iam/policies" // Import the new package
	sharpIam_repo "sharpnr-engine/core/sharp_iam/repo"
	sharpIam_services "sharpnr-engine/core/sharp_iam/services"

	"gorm.io/gorm"
)

func bootstrapSharIAM(db *gorm.DB) *sharpIam_services.SharpIamService {
	// Call the loader function directly
	basePolicy, err := policies.GetBasePolicy()
	if err != nil {
		panic("Failed to read base policy: " + err.Error())
	}

	engine := sharpIam_services.NewSharpIAM(basePolicy)
	sharpIamRepo := sharpIam_repo.NewSharpIamRepository(db)
	return sharpIam_services.NewSharpIamService(sharpIamRepo, engine)
}
