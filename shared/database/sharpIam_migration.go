package database

import (
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

	"gorm.io/gorm"
)

func migrateSharpIam(db *gorm.DB) error {
	return db.AutoMigrate(
		&sharp_iam_models.SharpIamRoles{},
		&sharp_iam_models.SharpIamEmployeeRoles{},
		&sharp_iam_models.SharpIamPolicies{},
	)
}
