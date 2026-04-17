package menu_controllers

import (
	"errors"
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *menuController) GetMenu(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	ownersIDRaw := ctx.Param("ownersID")
	parsedOwnersID, err := uuid.Parse(ownersIDRaw)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid ownersID format"))
		return
	}

	res, err := c.menuService.GetMenusOfOwners(ctx, orgID, emplID, parsedOwnersID)
	if err != nil {
		utils.Failure(ctx, http.StatusNotFound, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "", res)
}
