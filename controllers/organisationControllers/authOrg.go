package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos/auths"
	"github.com/sharpnr-inc/sharpnr-engine/models"
	services "github.com/sharpnr-inc/sharpnr-engine/services/orgnisation"
)

// RegisterOrganization handles registration of a new organization.
//
// This endpoint should be excluded from authentication middleware.
//
// Expected JSON Payload:
//
//	{
//	  "name": "Acme Corp",
//	  "email": "admin@acme.com",
//	  "password": "secret",
//	  ...
//	}
//
// Returns:
//   - 201 Created on success
//   - 400 Bad Request on validation failure
//   - 500 Internal Server Error on failure
//
// @Summary Register for organization Registration
// @Description This is Registration for the Organisation
// @Tags Organization
// @Accept  json
// @Produce  json
// @Param Register Organisation body dto.OrganisationRegisterDTO true "Login Credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /organization/register [post]
func RegisterOrganization(c *gin.Context) {
	var orgDTO dto.OrganisationRegisterDTO

	// Bind & validate JSON input based on `binding:"required"` tags
	if err := c.ShouldBindJSON(&orgDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	org := models.Organization{
		Name:                orgDTO.Name,
		OrganizationID:      orgDTO.OrganizationID,
		Address:             orgDTO.Address,
		Phone:               orgDTO.Phone,
		Email:               orgDTO.Email,
		IsPhysicalVerified:  boolPtrToVal(orgDTO.IsPhysicalVerified),
		ISRegistered:        boolPtrToVal(orgDTO.ISRegistered),
		CinNumber:           orgDTO.CinNumber,
		GstNumber:           orgDTO.GstNumber,
		OrganizationRegType: orgDTO.OrganizationRegType,
		Banker:              orgDTO.Banker,
		BankAccNumber:       orgDTO.BankAccNumber,
		IfscCode:            orgDTO.IfscCode,
		BankBranchName:      orgDTO.BankBranchName,
		Password:            orgDTO.Password,
		AdminName:           orgDTO.AdminName,
	}

	err := services.RegisterOrganizationService(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Organization registered successfully", "data": org})
}

// helper to convert *bool to bool, default false if nil
func boolPtrToVal(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// OrganizationRootLogin handles login for organizations.
//
// Accepts either email or username and a password.
// This endpoint should be excluded from authentication middleware.
//
// Expected JSON Payload:
//
//	{
//	  "email": "admin@acme.com",
//	  "password": "secret"
//	}
//
// Returns:
//   - 200 OK with JWT token
//   - 401 Unauthorized if credentials are invalid
//
// @Summary Login for organization root
// @Description Logs in an organization and returns JWT
// @Tags Organization
// @Accept  json
// @Produce  json
// @Param login body dto.LoginRequest true "Login Credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /organization/login [post]
func OrganizationRootLogin(c *gin.Context) {
	// Implementation here
}
