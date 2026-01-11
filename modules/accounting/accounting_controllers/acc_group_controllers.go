package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) GetAccountBookGroups(ctx *gin.Context) {
	// 1. Extract Org UUID (guaranteed by middleware, but still safe)
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

	// 2. Parse book ID from URL
	paramBookID := ctx.Param("bookId") // ⚠️ must match route param exactly
	bookID, err := uuid.Parse(paramBookID)
	if err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid accounting book ID",
			nil,
		)
		return
	}

	// 3. Call service
	resp, err := c.accSvc.GetAccountingGroupOfBook(
		ctx.Request.Context(),
		orgID,
		bookID,
	)
	if err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusInternalServerError,
			"Failed to fetch accounting groups",
			err,
		)
		return
	}

	// 4. Success response
	utils.SuccessResponse(
		ctx,
		http.StatusOK,
		"Accounting groups fetched successfully",
		resp,
	)
}
