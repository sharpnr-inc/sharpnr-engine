package accounting_controllers

import (
	"net/http"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) CreateAccountMapping(ctx *gin.Context) {
	// 1️⃣ Extract Org UUID from context
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

	// 2️⃣ Bind request body
	var req accountingdtos.CreateAccountMappingRequest_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid request payload",
			err,
		)
		return
	}

	// 3️⃣ Call service
	err := c.accSvc.CreateAccountMapping(ctx, orgID, req)
	if err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusInternalServerError,
			"Failed to create account mapping",
			err,
		)
		return
	}

	utils.SuccessResponse(
		ctx,
		http.StatusCreated,
		"Account mapping created successfully",
		"",
	)
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
