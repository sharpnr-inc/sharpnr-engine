package accounting_controllers

import (
	"sharpnr-engine/shared/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *accountingController) GetTransactions(ctx *gin.Context) {
	orgID := ctx.MustGet("OrgUUID").(uuid.UUID)

	bookID, err := uuid.Parse(ctx.Query("book_id"))
	if err != nil {
		utils.ErrorResponse(ctx, 400, "invalid book_id", nil)
	}

	from, _ := time.Parse("2006-01-02", ctx.DefaultQuery("from", time.Now().AddDate(0, -1, 0).Format("2006-01-02")))
	to, _ := time.Parse("2006-01-02", ctx.DefaultQuery("to", time.Now().Format("2006-01-02")))

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	data, total, err := c.accSvc.GetTransactions(
		ctx,
		orgID,
		bookID,
		from,
		to,
		limit,
		offset,
	)
	if err != nil {
		utils.ErrorResponse(ctx, 500, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, 200, "Transactions fetched", gin.H{
		"items": data,
		"total": total,
	})

}
