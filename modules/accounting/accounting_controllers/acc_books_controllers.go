package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) GetAccountingBooks(ctx *gin.Context) {
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

	// 2. Call service layer
	resp, err := c.accSvc.GetAccountingBooksOfOrg(
		ctx.Request.Context(),
		orgID,
	)
	if err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusInternalServerError,
			"Failed to fetch accounting books",
			err,
		)
		return
	}

	// 3. Success response (envelope)
	utils.SuccessResponse(
		ctx,
		http.StatusOK,
		"Accounting books fetched successfully",
		resp,
	)
}
