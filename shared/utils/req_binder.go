package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindJSON helper
func BindJSON[T any](ctx *gin.Context) (*T, bool) {
	var req T

	if err := ctx.ShouldBindJSON(&req); err != nil {
		Failure(ctx, http.StatusBadRequest, "Invalid Request Body", err)
		return nil, false
	}

	return &req, true
}
