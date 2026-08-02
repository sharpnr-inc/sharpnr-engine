package menu_controllers

import (
	"net/http"
	hospitality_menu_contracts "sharpnr-engine/business_ops/hospitality/menu/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *menuController) AddPreferredKitchenForMenuGroup(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownerID, ok := c.resolveMenuOwnersID(ctx)
	if !ok {
		return
	}

	var req hospitality_menu_contracts.SetKitchenToMenuGroupInput
	ok = c.resolveBody(ctx, &req) // ✅ pointer
	if !ok {
		return
	}

	if err := c.menuService.SetKitchenToMenuGroup(
		ctx.Request.Context(),
		orgID, emplID, ownerID,
		&req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen All set")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *menuController) AddPreferredKitchenForMenuCategory(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownerID, ok := c.resolveMenuOwnersID(ctx)
	if !ok {
		return
	}

	var req hospitality_menu_contracts.SetKitchenToMenuCategoryInput
	ok = c.resolveBody(ctx, &req) // ✅ pointer
	if !ok {
		return
	}

	if err := c.menuService.SetKitchenToMenuCategory(
		ctx.Request.Context(),
		orgID, emplID, ownerID,
		&req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen All set")

}

/* ---------------------------------- <***> --------------------------------- */
func (c *menuController) AddPreferredKitchenForMenuSubCategory(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownerID, ok := c.resolveMenuOwnersID(ctx)
	if !ok {
		return
	}

	var req hospitality_menu_contracts.SetKitchenToMenuSubCategoryInput
	ok = c.resolveBody(ctx, &req) // ✅ pointer
	if !ok {
		return
	}

	if err := c.menuService.SetKitchenToMenuSubCategory(
		ctx.Request.Context(),
		orgID, emplID, ownerID,
		&req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen All set")

}

/* ---------------------------------- <***> --------------------------------- */
func (c *menuController) AddPreferredKitchenForMenuItem(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownerID, ok := c.resolveMenuOwnersID(ctx)
	if !ok {
		return
	}

	var req hospitality_menu_contracts.SetKitchenToMenuItemInput
	ok = c.resolveBody(ctx, &req) // ✅ pointer
	if !ok {
		return
	}

	if err := c.menuService.SetKitchenToMenuItem(
		ctx.Request.Context(),
		orgID, emplID, ownerID,
		&req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen All set")
}
