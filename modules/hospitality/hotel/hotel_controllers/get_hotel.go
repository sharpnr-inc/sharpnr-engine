package hotelcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *HotelController) GetHotel(ctx *gin.Context) {
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

	// ✅ Call service
	res, err := s.hotelService.GetHotelDetails(ctx, orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)

}
