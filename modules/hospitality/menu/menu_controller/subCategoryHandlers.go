package menucontroller

import (
	"github.com/gin-gonic/gin"
)

func (c *menuController) CreateSubCategoryInCategory(ctx *gin.Context) {

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

	// // Bind Body
	// var req hospitalitydtos.Menu_CreateSubCategory_REQ
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid request payload",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// // Call service
	// err := c.menuService.CreateSubCategory(ctx, &req, orgID)
	// if err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Failed to create subcategory",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// // Success Response
	// res := dto.Response[any]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Subcategory created successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, res)
}

func (c *menuController) GetSubCategoryById(ctx *gin.Context) {

}
