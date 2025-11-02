package hotel

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Hospitalityhotelroutes(v1 *gin.RouterGroup, db *gorm.DB) {
	hotel_routes := v1.Group("/hotel")

	hotel_routes.GET("/get-all-rooms")
	hotel_routes.POST("/create-rooms")
	hotel_routes.POST("/book-room")

}
