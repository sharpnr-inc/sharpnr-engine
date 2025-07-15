package orgcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/sharpnr-inc/sharpnr-engine/core/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
)

type OrganisationController struct {
	service interfaces.CoreOrganisationService
}

func NewOrganisationController(service interfaces.CoreOrganisationService) *OrganisationController {
	return &OrganisationController{service: service}
}

func (c *OrganisationController) Register(ctx *gin.Context) {
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
		"Organisation": respDTO,
		"token":        token,
	})
}
