package menucontroller

import (
	"github.com/gin-gonic/gin"
)

func (c *menuController) CreateMenuGroup(ctx *gin.Context) {

}

func (c *menuController) GetAllMenuGroup(ctx *gin.Context) {
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 	return
	// }
	// // Type assert to UUID type (if you stored it as uuid.UUID)
	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 	return
	// }

	// // Call service → Already DTO
	// response, err := c.menuService.GetAllMenuGroups(ctx, orgID)
	// if err != nil {
	// 	res := dto.Response[any]{
	// 		StatusCode: http.StatusOK,
	// 		Message:    "There was an error Fetching the menu group",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, res)
	// 	return
	// }

	// // Success response
	// res := dto.Response[[]hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Menu groups fetched successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      &response,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, res)

}
