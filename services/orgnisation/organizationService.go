package services

import (
	"errors"

	"github.com/sharpnr-inc/sharpnr-engine/database"
	"github.com/sharpnr-inc/sharpnr-engine/models"
	"github.com/sharpnr-inc/sharpnr-engine/repositories"
	"golang.org/x/crypto/bcrypt"
)

func RegisterOrganizationService(org *models.Organization) error {
	// Basic validation already handled by Gin binding, but extra check
	if org.Name == "" || org.OrganizationID == "" || org.Phone == 0 || org.Password == "" {
		return errors.New("missing required fields")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(org.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	org.Password = string(hashedPassword)

	return repositories.CreateOrganization(database.DB, org)
}
