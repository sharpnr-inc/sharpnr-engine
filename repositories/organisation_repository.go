package repositories

import (
	"log"

	"github.com/sharpnr-inc/sharpnr-engine/models"
	"gorm.io/gorm"
)

func CreateOrganization(db *gorm.DB, org *models.Organization) error {
	if err := db.Create(org).Error; err != nil {
		log.Printf("❌ Failed to create organization: %v", err)
		return err
	}
	return nil
}
