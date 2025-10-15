package sharphrcontrollers

import (
	"fmt"
	"net/http"
	dto "sharpnr-engine/core/dtos"
	sharphrdtos "sharpnr-engine/modules/base/hr/sharphr_dtos"
	"sharpnr-engine/shared/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

type AddEmployee_RESPONSE_DTO struct {
	Message  string                   `json:"message"`
	Employee sharphrdtos.EmployeeData `json:"employee"`
}

func (c *sharpHrControllers) AddEmployee(ctx *gin.Context) {
	type AddEmployee_REQ_DTO struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Role  string `json:"role" binding:"required"`
	}

	var req AddEmployee_REQ_DTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// dummy employee data
	newEmployee := sharphrdtos.EmployeeData{
		ID:    "emp-" + fmt.Sprint(time.Now().UnixNano()),
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	responseData := AddEmployee_RESPONSE_DTO{
		Message:  "Employee added successfully",
		Employee: newEmployee,
	}

	successRes := dto.Response[AddEmployee_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "Employee added successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &responseData, // ✅ pointer matches *T
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, successRes)
}

func (c *sharpHrControllers) ListAllEmployee(ctx *gin.Context) {
	// For now, we just simulate dummy employee data
	dummyEmployees := []sharphrdtos.EmployeeData{
		{
			ID:    "1a2b3c4d-0001",
			Name:  "Alice Johnson",
			Email: "alice.johnson@example.com",
			Role:  "Manager",
		},
		{
			ID:    "1a2b3c4d-0002",
			Name:  "Bob Smith",
			Email: "bob.smith@example.com",
			Role:  "Engineer",
		},
		{
			ID:    "1a2b3c4d-0003",
			Name:  "Charlie Davis",
			Email: "charlie.davis@example.com",
			Role:  "HR Executive",
		},
	}

	// Wrap it inside your response DTO
	responseData := sharphrdtos.ListAllEmployee_RESPONSE_DTO{
		EmployeeList: dummyEmployees,
	}

	// Send a consistent response
	successRes := dto.Response[sharphrdtos.ListAllEmployee_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "Employee list fetched successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &responseData,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, successRes)
}

func (c *sharpHrControllers) GetEmployeeDetails(ctx *gin.Context) {
	// Step 1: Extract employeeId from URL param
	employeeID := ctx.Param("employeeId")
	if employeeID == "" {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Employee ID is required",
			Error: dto.ErrorResponse{
				Message: "Missing employee ID parameter",
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		})
		return
	}

	orgUUID, _ := ctx.Get("OrgUUID")

	// Step 3: Create dummy employee data
	dummyEmployee := sharphrdtos.GetEmployeeDetails_RESPONSE_DTO{
		ID:           employeeID,
		OrgID:        orgUUID,
		FirstName:    "Kashif",
		LastName:     "Nawaz",
		Email:        "kashif.nawaz@example.com",
		Phone:        "+91 9876543210",
		Designation:  "Software Engineer",
		Department:   "R&D",
		JoiningDate:  "2024-01-15",
		AccessLevel:  "Admin",
		IsActive:     true,
		CreatedAt:    time.Now().AddDate(0, -6, 0),
		LastModified: time.Now(),
	}

	// Step 5: Return success response
	ctx.JSON(http.StatusOK, dto.Response[sharphrdtos.GetEmployeeDetails_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &dummyEmployee,
		Timestamp: time.Now(),
	})
}

func (c *sharpHrControllers) EmployeeExists(ctx *gin.Context) {
	var req sharphrdtos.CheckIfEmployeeExist_REQ_DTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "Something Went Wrong",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			}, Data: nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	orgSharpnrID := ctx.GetString(middlewares.OrgIDContextKey)
	fmt.Println(orgSharpnrID)
	orgUUID, _ := ctx.Get("OrgUUID")

	fmt.Println(orgUUID)

	res, err := c.service.GetEmployeeDetailByEmail(orgSharpnrID, req.EmployeeEmail, orgUUID)
	if err != nil {
		errRes := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "Some Went Wrong",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusOK, errRes)
		return
	}

	successRes := dto.Response[sharphrdtos.CheckIfEmployeeExist_RESPONSE_DTO]{
		StatusCode: http.StatusOK,
		Message:    "Hey Contrats",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      res,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, successRes)

}
