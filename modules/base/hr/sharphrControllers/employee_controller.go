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

func (c *sharpHrControllers) AddEmployee(ctx *gin.Context) {
}

func (c *sharpHrControllers) ListAllEmployee(ctx *gin.Context) {

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
