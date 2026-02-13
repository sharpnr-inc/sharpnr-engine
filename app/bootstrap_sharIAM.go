package app

import (
	"embed"
	sharpIam_repo "sharpnr-engine/core/sharp_iam/repo"
	sharpIam_services "sharpnr-engine/core/sharp_iam/services"

	"gorm.io/gorm"
)

var policyFS embed.FS

func bootstrapSharIAM(db *gorm.DB) *sharpIam_services.SharpIamService {

	basePolicy, err := policyFS.ReadFile("policies/requirements.cedar")
	if err != nil {
		panic("Failed to read base policy: " + err.Error())
	}

	// Pass the policy bytes to the engine
	engine := sharpIam_services.NewSharpIAM(basePolicy)
	sharpIamRepo := sharpIam_repo.NewSharpIamRepository(db)
	return sharpIam_services.NewSharpIamService(sharpIamRepo, engine)
}
