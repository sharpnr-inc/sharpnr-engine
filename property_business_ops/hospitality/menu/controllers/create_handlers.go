package menu_controllers

import (
	"errors"
	"net/http"
	hospitality_menu_contracts "sharpnr-engine/property_business_ops/hospitality/menu/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddMenu handles the root menu creation (Recursive or hollow)
func (c *menuController) AddMenu(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.CreateMenuInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.CreateMenu(ctx, orgID, emplID, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}

func (c *menuController) AddCategoryInMenu(ctx *gin.Context) {

	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownersID := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid restaurantID format"))
		return
	}

	var req *hospitality_menu_contracts.AddMenuCategoryInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddCategoryInMenu(ctx, req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Success")

}

// AddSubCategory adds a sub-category to an existing category
func (c *menuController) AddSubCategory(ctx *gin.Context) {
	_, ok := c.validateContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.AddSubCategoryInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddSubCategoryInCategory(ctx, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}

// AddMenuItems adds an item to a sub-category
func (c *menuController) AddMenuItems(ctx *gin.Context) {
	_, ok := c.validateContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.AddMenuItemInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddMenuItemInSubCategory(ctx, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}

// AddVariantsGroup adds a variant group (like 'Size' or 'Add-ons') to an item
func (c *menuController) AddVariantsGroup(ctx *gin.Context) {
	_, ok := c.validateContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.AddVariantGroupInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddVariantGroup(ctx, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}

// AddVariants adds specific variants to a variant group
func (c *menuController) AddVariants(ctx *gin.Context) {
	_, ok := c.validateContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.AddVariantInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddVariantsInVariantsGroup(ctx, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}

// AddAddons adds an addon to a menu item
func (c *menuController) AddAddons(ctx *gin.Context) {
	_, ok := c.validateContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	_, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid ownersID format"))
		return
	}

	var req hospitality_menu_contracts.AddAddOnInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if _, err := c.menuService.AddAddOnInMenu(ctx, &req); err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "success")
}
