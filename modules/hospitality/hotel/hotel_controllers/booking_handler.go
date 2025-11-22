package hotelcontrollers

import (
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *HotelController) GetBookingDetails(ctx *gin.Context) {
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	roomId := ctx.Param("roomId")
	roomUUID, err := uuid.Parse(roomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room UUID"})
		return
	}

	booking, err := c.hotelService.GetRoomBookingDetails(ctx, roomUUID, orgID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"booking": booking,
	})
}

func (c *HotelController) CreateBooking(ctx *gin.Context) {

	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	// Type assert to UUID type (if you stored it as uuid.UUID)
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	var req hospitalitydtos.Hotel_Room_booking_Req

	// Bind JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid or missing fields",
			"details": err.Error(),
		})
		return
	}

	// Call service
	if err := c.hotelService.CreateBookingForRoom(ctx, orgID, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Room booked successfully",
	})

}
