package menu_controllers

import (
	"errors"
	"net/http"
	menuservice "sharpnr-engine/property_resources/hospitality_resources/menu/service"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type menuController struct {
	menuService menuservice.MenuService
}

func NewMenuController(menuService menuservice.MenuService) *menuController {
	return &menuController{
		menuService: menuService,
	}
}

// Helper to validate basic auth and org context to avoid repetition
func (c *menuController) validateContext(ctx *gin.Context) (uuid.UUID, bool) {
	if _, exist := ctx.Get(webMiddleware.UserIDContextKey); !exist {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, false
	}

	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, false
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid type"))
		return uuid.Nil, false
	}

	return orgID, true
}
