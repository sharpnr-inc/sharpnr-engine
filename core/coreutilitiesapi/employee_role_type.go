package coreutilitiesapi

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEmployeeRolesTypes(c *gin.Context) {
	employeeRoles := []constants.EmployeeRoles{
		constants.Admin,
		constants.Manager,
	}

	// ✅ Wrap inside key
	wrappedData := map[string]any{
		"EMPLOYEE_ROLE_TYPE": employeeRoles,
	}

	// ✅ Use standard response DTO
	res := dto.Response[map[string]any]{
		StatusCode: http.StatusOK,
		Message:    "List of All Employee Role Types",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &wrappedData, // pointer because Data is *T
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, res)
}
