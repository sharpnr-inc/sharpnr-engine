package accounting_controllers

import (
	"fmt"
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) GetAccounts(ctx *gin.Context) {
	orgUUIDVal, exists := ctx.Get("OrgUUID")
	if !exists {
		utils.ErrorResponse(
			ctx,
			http.StatusUnauthorized,
			"Organization UUID missing",
			nil,
		)
		return
	}

	orgID, ok := orgUUIDVal.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid organization UUID",
			nil,
		)
		return
	}

	paramsBookId := ctx.Param("bookID")

	bookID, err := uuid.Parse(paramsBookId)
	if err != nil {
		utils.ErrorResponse(ctx,
			http.StatusBadRequest,
			"Invalid accounting bookID",
			nil)
		return
	}
	fmt.Println(bookID)
	fmt.Println(orgID)
}
