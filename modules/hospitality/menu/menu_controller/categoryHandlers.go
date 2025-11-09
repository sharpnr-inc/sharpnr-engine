package menucontroller

// This is Category Handlers
import (
	"net/http"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *menuController) CreateCategoryInMenu(ctx *gin.Context) {

}

func (c *menuController) GetAllCategoryInMenu(ctx *gin.Context) {

}

func (c *menuController) GetCategoriesByGroup(ctx *gin.Context) {

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
