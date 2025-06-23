package services

import (
	"errors"

	"github.com/sharpnr-inc/sharpnr-engine/models"
	"github.com/sharpnr-inc/sharpnr-engine/repositories"
)

func RegisterOrganizationService(org *models.Organization) error {
	if org.Name == "" || org.OrganizationID == "" {
		return errors.New("organization name and ID are required")
	}
	return repositories.CreateOrganization(org)
}
