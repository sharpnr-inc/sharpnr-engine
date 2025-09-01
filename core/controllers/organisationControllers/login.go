package orgcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *OrganisationController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"redirectUrl": "sharpnrapp://auth-callback?token=12345",
		"token":       "12345",
	})
}
