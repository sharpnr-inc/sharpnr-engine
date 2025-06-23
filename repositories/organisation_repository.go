package repositories

import (
	"github.com/sharpnr-inc/sharpnr-engine/database"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

func CreateOrganization(org *models.Organization) error {
	return database.DB.Create(org).Error
}
