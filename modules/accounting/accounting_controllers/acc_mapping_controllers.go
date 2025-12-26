package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) CreateAccountMapping(ctx *gin.Context) {
	// it will takes the Request and create the mapping for the Accounting book
	panic("unimplemented")
}

func (c *accountingController) GetAllAccountMapping(ctx *gin.Context) {
	// 1. Extract Org UUID from context (set by auth middleware)
	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		utils.ErrorResponse(
			ctx,
			http.StatusUnauthorized,
			"Organization UUID missing",
			nil,
		)
		return
	}

	orgID, ok := orgUUIDVal.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid organization UUID",
			nil,
		)
		return
	}

	data, err := c.accSvc.GetAccountMapping(ctx, orgID)
	if err != nil {
		utils.ErrorResponse(ctx, 200, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, 200, "Account mappings fetched", data)

}
