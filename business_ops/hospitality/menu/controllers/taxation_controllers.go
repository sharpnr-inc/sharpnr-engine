package menu_controllers

import (
	"net/http"
	hospitality_menu_contracts "sharpnr-engine/business_ops/hospitality/menu/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *menuController) AddTaxCategoryInMenuGroup(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownersID, ok := c.resolveMenuOwnersID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[hospitality_menu_contracts.MenuGroupTaxCategoryRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.menuService.AddTaxCategoryInMenuGroup(ctx.Request.Context(), orgID, emplID, ownersID, *req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong During Setting the Tax Category", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Tax Category Added to Menu Group")
}
