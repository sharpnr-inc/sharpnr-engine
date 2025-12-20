package menucontroller

// This is Category Handlers
import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCategoryInMenu godoc
// @Summary Create Category In Menu
// @Description Create Category linked to the authenticated organization UUID.
// @Tags Hospitality / Menu
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "Category created successfully"
// @Failure 400 {object} map[string]string "Invalid organization UUID format"
// @Router /hospitality-menu/categories [post]
func (c *menuController) CreateCategoryInMenu(ctx *gin.Context) {

	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	var req hospitalitydtos.Menu_CreateCategoryINMenuGroup_REQ
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call service
	err := c.menuService.CreateCategory(ctx, &req, orgID)
	if err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create category",
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

	// Success response
	res := dto.Response[any]{
		StatusCode: http.StatusOK,
		Message:    "Category created successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *menuController) GetAllCategoryInMenu(ctx *gin.Context) {

}

func (c *menuController) GetCategoriesByGroup(ctx *gin.Context) {
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// Get MenuGroupId from params
	groupIDStr := ctx.Param("id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group id"})
		return
	}

	// Call service
	result, err := c.menuService.GetCategoriesByGroup(ctx, orgID, groupID)
	if err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to fetch categories",
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

	res := dto.Response[[]hospitalitydtos.Menu_Category_RESPONSE]{
		StatusCode: http.StatusOK,
		Message:    "Categories fetched successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &result,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *menuController) GetCategoryInMenuByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	category, err := c.menuService.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *menuController) UpdateCategoryInMenu(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	var category hospitalitymodels.MenuCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.ID = id

	if err := c.menuService.UpdateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (c *menuController) DeleteCategoryInMenu(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	if err := c.menuService.DeleteCategory(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
