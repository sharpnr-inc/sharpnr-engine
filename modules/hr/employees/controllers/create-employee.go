package hr_employee_controllers

import (
	"net/http"
	sharphrdtos "sharpnr-engine/modules/hr/hr_dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *hrEmployeeController) AddEmployeeInOrg(ctx *gin.Context) {

	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Organization UUID missing", nil)
		return
	}

	orgID := orgUUIDVal.(uuid.UUID)

	var req sharphrdtos.CreateEmployeeInOrganization_RequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	if err := c.hrEmpSvc.CreateNewEmployee(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Done": "Mail Send"})
		return
	}

}
