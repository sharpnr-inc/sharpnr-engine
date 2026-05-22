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

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetLegalEntityDetailsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// we will take out the legalEntityID from ID
	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	response, err := c.orchestService.GetDetailsOfLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "some thing went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Details Fetched successfully", response)
}

/* ---------------------- Add Contacts in Legal Entity ---------------------- */

func (c *OrchestratorController) AddContactsInLegalEntityHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// we will take out the legalEntityID from ID
	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	req, ok := utils.BindJSON[legal_entity_contracts.CreateLegalEntityContactsRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddContactsInLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Contact Added Successfully")
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddAddressesInLegalEntityHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// we will take out the legalEntityID from ID
	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	req, ok := utils.BindJSON[legal_entity_contracts.CreateLegalEntityAddressRequest](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddAddressInLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Address Added")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddBankAccountDetailsInLegalEntityHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// we will take out the legalEntityID from ID
	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	req, ok := utils.BindJSON[legal_entity_contracts.AddBankAccountInLegalEntityRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddBankAccountDetailsInLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Bank Account Added")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddIdentifiersInLegalEntityHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// we will take out the legalEntityID from ID
	legalEntityID, err := c.getLegalEntityIDFromParams(ctx)
	if err != nil {
		return
	}

	req, ok := utils.BindJSON[legal_entity_contracts.AddIdentifiersInLegalEntityRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddIdentifiersInLegalEntity(ctx.Request.Context(), orgID, emplID, legalEntityID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Identifiers Added")

}

/* ---------------------------------- <***> --------------------------------- */

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

// GetLegalEntityTaxTypes handles GET /utils/tax_types
func (c *OrchestratorController) GetLegalEntityTaxTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllTaxTypes()
	response := make([]legal_entity_constants.LegalEntityTaxTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityTaxTypeResponse{
			ID:          string(t),
			Label:       t.Label(),
			Description: t.Description(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity tax types fetched successfully", response)
}

// GetLegalEntityTaxComponentTypes handles GET /utils/tax-components_types
func (c *OrchestratorController) GetLegalEntityTaxComponentTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllTaxComponentTypes()
	response := make([]legal_entity_constants.LegalEntityTaxComponentTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityTaxComponentTypeResponse{
			ID:          string(t),
			Label:       t.Label(),
			Description: t.Description(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity tax component types fetched successfully", response)
}

// GetLegalEntityBankAccountTypes handles GET /utils/bank-acc_types
func (c *OrchestratorController) GetLegalEntityBankAccountTypes(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	types := legal_entity_constants.AllBankAccountTypes()
	response := make([]legal_entity_constants.LegalEntityBankAccountTypeResponse, len(types))
	for i, t := range types {
		response[i] = legal_entity_constants.LegalEntityBankAccountTypeResponse{
			ID:          string(t),
			Label:       t.Label(),
			Description: t.Description(),
		}
	}
	utils.Success(ctx, http.StatusOK, "Legal entity bank account types fetched successfully", response)
}
