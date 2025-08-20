package orgcontrollers

import (
	"net/http"

	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/interfaces"

	"github.com/gin-gonic/gin"
)

type OrganisationController struct {
	service interfaces.OrgRegistrationService
}

func RegisterOrgController(service interfaces.OrgRegistrationService) *OrganisationController {
	return &OrganisationController{service: service}
}

func (c *OrganisationController) InitOrgReg(ctx *gin.Context) {
	var req dto.OrgInitialRegistration_Request_Dto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resDTO, _, err := c.service.InitiateOrgRegistration(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success response
	ctx.JSON(http.StatusOK, gin.H{
		"data": resDTO,
	})

}
