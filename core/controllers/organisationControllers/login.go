package orgcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *OrganisationController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"redirectUrl": "sharpnrapp://dashboard",
		"token":       "12345",
	})
}
