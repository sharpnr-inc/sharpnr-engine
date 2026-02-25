package hospitality_utilities

import "github.com/gin-gonic/gin"

func RegisterHospitalityUtilityRoutes(v1 *gin.RouterGroup) {
	ut := v1.Group("/utils")
	ut.GET("/:domain", GetHospitalityDomainStencil)

}
