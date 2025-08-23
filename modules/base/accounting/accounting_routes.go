package accounting

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AccountingRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

}
