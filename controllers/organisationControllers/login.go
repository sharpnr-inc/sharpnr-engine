package orgcontrollers

import (
	"net/http"

	dto "sharpnr-engine/core/dtos"

	"github.com/gin-gonic/gin"
)

func (c *OrganisationController) Login(ctx *gin.Context) {
	var req dto.OrgAuthRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
