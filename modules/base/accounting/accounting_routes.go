package accounting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	v1.POST("/acc/creat-invoice")
	v1.POST("/acc/list-invoice")
	v1.POST("/acc/add-expense")
	v1.POST("/acc/ping", PingHandler)
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
