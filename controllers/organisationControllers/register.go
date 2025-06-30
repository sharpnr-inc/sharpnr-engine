package orgcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/interfaces"
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
// @Param Register Organisation body dto.OrganizationRegistrationRequestDTO true "Registration Details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /organization/register [post]
type OrganizationController struct {
	service interfaces.OrganizationService
}

func NewOrganizationController(service interfaces.OrganizationService) *OrganizationController {
	return &OrganizationController{service: service}
}

func (c *OrganizationController) Register(ctx *gin.Context) {
	var req dto.OrganizationRegistrationRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respDTO, token, err := c.service.RegisterOrganization(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"organization": respDTO,
		"token":        token,
	})
}
