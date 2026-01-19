package coreutilitiesapi

import (
	"net/http"

	"sharpnr-engine/core/constants"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func GetEmployeeRolesTypes(c *gin.Context) {
	employeeRoles := []constants.EmployeeRoles{
		constants.Admin,
		constants.Manager,
	}

	// Wrap inside key (frontend-friendly shape)
	wrappedData := map[string]any{
		"EMPLOYEE_ROLE_TYPE": employeeRoles,
	}

	utils.Success(
		c,
		http.StatusOK,
		"List of All Employee Role Types",
		wrappedData,
	)
}
