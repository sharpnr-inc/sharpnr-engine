package orchestrator_controllers

import (
	"net/http"
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*           All controllers legal entity management will be defined          */
/* -------------------------------------------------------------------------- */

/* -------------------------- list all Legal Entity ------------------------- */

func (c *OrchestratorController) ListAllLegalEntity(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetAllLegalEntity(ctx, orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Legal Entities retrieved successfully", res)
}

/* ---------------------------- Add Legal Entity ---------------------------- */

func (c *OrchestratorController) AddLegalEntity(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	var req *legal_entity_contracts.NewLegalEntityRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}

	if err := c.orchestService.CreateLegalEntity(ctx, orgID, emplID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Legal Entity Created successfully")
}

/* -------------------------------------------------------------------------- */
/*                        Legal entitys utils endpoints                       */
/* -------------------------------------------------------------------------- */

func (c *OrchestratorController) GetEntityTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllBusinessRegistrationTypes()
	response := make([]legal_entity_constants.EntityTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.EntityTypeResponse{
			ID:    string(t),
			Label: t.Label(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity types fetched successfully ", response)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetGSTRegistrationTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllGSTRegistrationTypes()
	response := make([]legal_entity_constants.GSTRegistrationTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.GSTRegistrationTypeResponse{
			ID:    string(t),
			Label: t.Label(),
		}
	}
	utils.Success(ctx, http.StatusOK, "GST Registration types fetched successfully ", response)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetLegalEntityAddressTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllAddressTypes()
	response := make([]legal_entity_constants.LegalEntityAddressTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityAddressTypeResponse{
			ID:    string(t),
			Label: t.Label(),
		}
	}
	utils.Success(ctx, http.StatusOK, "GST Types Fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */

// GetLegalEntityLicenseTypes handles GET /utils/license_types
func (c *OrchestratorController) GetLegalEntityLicenseTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllLicenseTypes()
	response := make([]legal_entity_constants.LegalEntityLicenseTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityLicenseTypeResponse{
			ID:    string(t),
			Label: t.Label(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity license types fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */

// GetLegalEntityContactTypes handles GET /utils/contact_types
func (c *OrchestratorController) GetLegalEntityContactTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllContactTypes()
	response := make([]legal_entity_constants.LegalEntityContactTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityContactTypeResponse{
			ID:    string(t),
			Label: t.Label(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity contact types fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */

// GetLegalEntityIdentifierTypes handles GET /utils/identifier_types
func (c *OrchestratorController) GetLegalEntityIdentifierTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllIdentifierTypes()
	response := make([]legal_entity_constants.LegalEntityIdentifierTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityIdentifierTypeResponse{
			ID:          string(t),
			Label:       t.Label(),
			Description: t.Description(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity identifier types fetched successfully", response)
}
