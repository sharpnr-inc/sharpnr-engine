package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *HotelController) CreateHotelGuest(ctx *gin.Context) {
	panic("unimplemented")
}

func (s *HotelController) SearchGuests(ctx *gin.Context) {
	// ✅ Get OrgUUID from context (set by middleware)
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

	// 2️⃣ Get search query (query param, not path)
	searchQry := strings.TrimSpace(ctx.Query("q"))
	if searchQry == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	guests, err := s.hotelService.SearchGuests(ctx, orgID, searchQry)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := dto.Response[[]hoteldtos.HotelGuest_dto]{
		StatusCode: http.StatusOK,
		Message:    "Guest In Hotels",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &guests,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, response)
}
