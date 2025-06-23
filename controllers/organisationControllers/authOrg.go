package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Param login body dto.LoginRequest true "Login Credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /organization/register [post]
func RegisterOrganization(c *gin.Context) {
	var org models.Organization

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	err := services.RegisterOrganizationService(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Organization registered successfully", "data": org})
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
