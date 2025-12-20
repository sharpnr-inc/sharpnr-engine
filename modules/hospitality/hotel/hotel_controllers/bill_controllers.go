package hotelcontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *HotelController) PreviewBill(ctx *gin.Context) {

	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Organization UUID missing", nil)
		return
	}

	orgID := orgUUIDVal.(uuid.UUID)

	bookingIDStr := ctx.Param("bookingId")
	bookingID, err := uuid.Parse(bookingIDStr)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid booking ID", nil)
		return
	}

	preview, err := c.hotelService.PreviewBill(ctx, orgID, bookingID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Bill preview", preview)

}
