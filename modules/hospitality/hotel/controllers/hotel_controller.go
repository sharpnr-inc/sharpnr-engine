package hotelcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HotelController) CreateRoom(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}
