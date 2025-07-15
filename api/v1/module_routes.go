package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	// This is where you would initialize module-specific controllers and routes.
	// For example:

	// --- Hotel Module Routes ---
	// hotelService := hotel.NewService(...)
	// hotelController := hotel.NewController(hotelService)
	// v1.POST("/hotel/bookings", hotelController.CreateBooking)

	// --- Retail Module Routes ---
	// retailService := retail.NewService(...)
	// retailController := retail.NewController(retailService)
	// v1.POST("/retail/sales", retailController.RecordSale)
}