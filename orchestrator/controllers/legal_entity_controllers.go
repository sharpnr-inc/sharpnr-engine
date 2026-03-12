package orchestrator_controllers

import (
	"net/http"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*           All controllers legal entity management will be defined          */
/* -------------------------------------------------------------------------- */

/* ---------------------------- Add Legal Entity ---------------------------- */

func (c *OrchestratorController) AddLegalEntity(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	var req *orchestrator_dtos.AddLegalEntity_REQ_DTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.CreateLegalEntity(ctx, orgID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Legal Entity Created successfully")
}

/* -------------------------- list all Legal Entity ------------------------- */

func (c *OrchestratorController) ListAllLegalEntity(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetAllLegalEntity(ctx, orgID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Legal Entities retrieved successfully", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddTaxRegistrationLegalEntity(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	req, ok := utils.BindJSON[legal_entity_contracts.CreateLegalEntityTaxRegistrationRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddTaxRegsitrationInLegalEntity(
		ctx.Request.Context(), orgID, emplID, legalEntityID, *req,
	); err != nil {
		utils.Failure(
			ctx, http.StatusInternalServerError,
			"something went Wrong",
			err,
		)
		return
	}

	utils.Success(ctx, http.StatusOK, "Tax Registration Added SuccessFully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) ListAllTaxRegistrationOfLegalEntity(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	res, err := c.orchestService.GetAllTaxRegistrationOfLegalEntity(
		ctx.Request.Context(),
		orgID, emplID, legalEntityID,
	)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Success", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddTaxCategoryInLegalEntity(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	request, ok := utils.BindJSON[legal_entity_contracts.CreateLegalEntityTaxCategory](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddTaxCategoryInLegalEntiy(
		ctx.Request.Context(),
		orgID, emplID, legalEntityID, *request,
	); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Added Successfully")
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) ListAllTaxCategoryOfLegalEntry(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	response, err := c.orchestService.ListAllTaxCategoryOfLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Fetched Successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */
