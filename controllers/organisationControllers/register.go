package orgcontrollers

import (
	"net/http"

	dto "sharpnr-engine/core/dtos"

	"github.com/gin-gonic/gin"
)

// InitOrgReg godoc
// @Summary      Register a new organization
// @Description  Start the registration process for a new organization
// @Tags         Registration
// @Accept       json
// @Produce      json
// @Param        request body dto.OrgInitialRegistration_Request_Dto true "Organization registration payload"
// @Success      200  {object} dto.OrgInitialRegistration_Response_Dto
// @Failure      400  {object} map[string]string
// @Router       /org/register [post]
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
