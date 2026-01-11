package accounting_controllers

import (
	"fmt"
	"net/http"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) GetAccounts(ctx *gin.Context) {
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

	paramsBookId := ctx.Param("bookID")

	bookID, err := uuid.Parse(paramsBookId)
	if err != nil {
		utils.ErrorResponse(ctx,
			http.StatusBadRequest,
			"Invalid accounting bookID",
			nil)
		return
	}
	fmt.Println(bookID)
	fmt.Println(orgID)
}

func (c *accountingController) CreateAccount(ctx *gin.Context) {
	orgID := ctx.MustGet("OrgUUID").(uuid.UUID)

	var req accountingdtos.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, 400, err.Error(), nil)
		return
	}

	if err := c.accSvc.CreateAccount(ctx, orgID, req); err != nil {
		utils.ErrorResponse(ctx, 400, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, 201, "Account created successfully", "")
}
