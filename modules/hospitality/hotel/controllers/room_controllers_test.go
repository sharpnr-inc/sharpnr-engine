package hotelcontrollers_test

import (
	hotelcontrollers "sharpnr-engine/modules/hospitality/hotel/hotel_controllers"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHotelController_CreateRoom(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		s hotelservices.HotelService
		// Named input parameters for target function.
		ctx *gin.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := hotelcontrollers.NewHotelController(tt.s)
			s.CreateRoom(tt.ctx)
		})
	}
}
