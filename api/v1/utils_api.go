package v1

import (
	"net/http"

	"sharpnr-engine/core/constants"

	"github.com/gin-gonic/gin"
)

// UtilsRoutes registers utility endpoints (for enums, static data, etc.)
func UtilsRoutes(rg *gin.RouterGroup) {
	utils := rg.Group("/utils")

	utils.GET("/org-types", GetOrgTypes)
	utils.GET("/org-reg-types", GetBusinessRegistrationType)
	utils.GET("/employee-role-types", GetEmployeeRolesTypes)
}

// GetOrgTypes returns the list of available organization types.
func GetOrgTypes(c *gin.Context) {
	orgTypes := []constants.OrgType{
		constants.Hotel,
		constants.Retail,
		constants.Salon,
		constants.Club,
		constants.Restaurant,
	}

	c.JSON(http.StatusOK, gin.H{"orgTypes": orgTypes})
}

func GetBusinessRegistrationType(c *gin.Context) {
	businessRegType := []constants.BusinessRegistrationType{
		constants.SoleProprietorship,
		constants.Partnership,
		constants.LimitedLiabilityPartnership,
		constants.PrivateLimitedCompany,
		constants.PublicLimitedCompany,
		constants.OnePersonCompany,
		constants.Section8Company,
		constants.CooperativeSociety,
		constants.Trust,
		constants.Society,
	}

	c.JSON(http.StatusOK, gin.H{
		"business_registration_types": businessRegType,
	})
}

func GetEmployeeRolesTypes(c *gin.Context) {
	employeeRoles := []constants.EmployeeRoles{
		constants.Admin,
		constants.Manager,
	}

	c.JSON(http.StatusOK, gin.H{"employee_roles": employeeRoles})
}
