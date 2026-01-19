package menucontroller

import (
	"github.com/gin-gonic/gin"
)

func (c *menuController) CreateMenuItem(ctx *gin.Context) {

	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 	return
	// }

	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 	return
	// }

	// var req hospitalitydtos.Menu_CreateMenuItem_REQ
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid request payload",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// // default served_by if not sent
	// if req.ServedBy == "" {
	// 	req.ServedBy = "KITCHEN"
	// }

	// response, err := c.menuService.CreateMenuItem(ctx, &req, orgID)
	// if err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Failed to create menu item",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// res := dto.Response[hospitalitydtos.Menu_CreateMenuItem_RES]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Menu item created successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      response,
	// 	Timestamp: time.Now(),
	// }
	// ctx.JSON(http.StatusOK, res)
}
