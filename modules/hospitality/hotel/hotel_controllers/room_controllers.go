package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *HotelController) CreateRoom(ctx *gin.Context) {
	var req hoteldtos.CreateRoomInFloor_REQ

	// ✅ Bind and validate request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// ✅ Get OrgUUID from context (set by middleware)
	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgUUID, ok := orgUUIDVal.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// ✅ Call service
	if err := s.hotelService.CreateRoomInHotel(ctx, orgUUID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Room Added successfully"})
}

// This is the Update Room Status
func (s *HotelController) UpdateRoomStatus(ctx *gin.Context) {
	// ✅ Get OrgUUID from context (set by middleware)
	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgUUID, ok := orgUUIDVal.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// -------------------------------------------------------
	// 2️⃣ Extract URL Params (hotelId, roomId, status)
	// -------------------------------------------------------
	hotelIdStr := ctx.Param("hotelId")
	roomIdStr := ctx.Param("roomId")
	status := ctx.Param("status") // no validation needed unless you want enum

	// -------------------------------------------------------
	// 3️⃣ Validate hotelId as UUID
	// -------------------------------------------------------
	hotelUUID, err := uuid.Parse(hotelIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotelId UUID"})
		return
	}

	// -------------------------------------------------------
	// 4️⃣ Validate roomId as UUID
	// -------------------------------------------------------
	roomUUID, err := uuid.Parse(roomIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roomId UUID"})
		return
	}

	if err := s.hotelService.UpdateRoomStatus(ctx, orgUUID, hotelUUID, roomUUID, status); err != nil {
		response := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "There was error updating",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusOK, response)
		return
	}

	// 6️⃣ SUCCESS RESPONSE
	response := dto.Response[any]{
		StatusCode: http.StatusOK,
		Message:    "Room status updated successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, response)

}
