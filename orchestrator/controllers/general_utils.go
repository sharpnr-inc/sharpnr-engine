package orchestrator_controllers

import (
	"net/http"
	general_infra_utils "sharpnr-engine/business_ops/general_infra/utils"
	healthcare_utilities "sharpnr-engine/business_ops/healthcare/utility"
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) BusinessOperationModulesHandler(ctx *gin.Context) {

	// Ignore the organization ID and employeeID
	// this is for Authentication purpose, so no loose enpoint is exposed
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	data := gin.H{
		"hospitality": hospitality_utilities.HospitalityModules,
		"healthcare":  healthcare_utilities.HealthcareModules,
	}

	utils.Success(ctx, http.StatusOK, "Business modules fetched successfully", data)
}

func (c *OrchestratorController) InfraModulesHandler(ctx *gin.Context) {

	// Ignore the organization ID and employeeID
	// this is for Authentication purpose, so no loose enpoint is exposed
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	data := gin.H{
		"general":     general_infra_utils.GeneralInfraModules,
		"hospitality": hospitality_utilities.HospitalityInfraModules,
		"healthcare":  healthcare_utilities.HealthcareInfraModules,
	}

	utils.Success(ctx, http.StatusOK, "Infra modules fetched successfully", data)
}
