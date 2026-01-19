package hotelcontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *HotelController) GetHotelFrontDesk(ctx *gin.Context) {
	// TODO we need
	// 	orgUUID, exists := ctx.Get("OrgUUID")
	// 	if !exists {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 		return
	// 	}
	// 	// Type assert to UUID type (if you stored it as uuid.UUID)
	// 	orgID, ok := orgUUID.(uuid.UUID)
	// 	if !ok {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 		return
	// 	}

	// 	res, err := c.hotelService.GetFrontDeskTerminal(ctx, orgID)
	// 	if err != nil {
	// 		response := dto.Response[any]{
	// 			StatusCode: http.StatusOK,
	// 			Message:    "There was an error fetching the Frontdesk Terminal",
	// 			Error: dto.ErrorResponse{
	// 				Message: err.Error(),
	// 				IsError: true,
	// 			},
	// 			Data:      nil,
	// 			Timestamp: time.Now(),
	// 		}
	// 		ctx.JSON(http.StatusOK, response)
	// 		return
	// 	}

	// 	response := dto.Response[hospitalitydtos.HotelBookingTerminal_Response]{
	// 		StatusCode: http.StatusOK,
	// 		Message:    "Booking Terminal Fetch successfully",
	// 		Error: dto.ErrorResponse{
	// 			Message: "",
	// 			IsError: false,
	// 		},
	// 		Data:      res,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusOK, response)

	// }

	//

	// Here we Need to get the Hotel Frontdesk
}

func (c *HotelController) GetHotelReception(ctx *gin.Context) {
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
	fmt.Println(orgID)
}
