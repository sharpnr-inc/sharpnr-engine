package hotelcontrollers

import (
	"net/http"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HotelController) CreateTaxRule(ctx *gin.Context) {
	// -------------------------
	// 1️⃣ Get Org UUID from context
	// -------------------------
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

	orgUUID, ok := orgUUIDVal.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid organization UUID format",
			nil,
		)
		return
	}

	// -------------------------
	// 2️⃣ Bind request body
	// -------------------------
	var req hoteldtos.CreateHotelTaxRuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	// -------------------------
	// 3️⃣ Call service
	// -------------------------
	if err := h.hotelService.CreateTaxRule(ctx, orgUUID, req); err != nil {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Failed to create tax rule",
			err,
		)
		return
	}

	// -------------------------
	// 4️⃣ Success response
	// -------------------------
	var data any = nil
	utils.SuccessResponse(
		ctx,
		http.StatusCreated,
		"Tax rule created successfully",
		data,
	)
}
