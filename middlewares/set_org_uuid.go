package middlewares

import (
	"net/http"

	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/database"
	"time"

	"github.com/gin-gonic/gin"
)

func SetOrgUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the SharpnrID from context (set by RequireTokenType)
		sharpnrIDVal, exists := ctx.Get(OrgIDContextKey)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Missing SharpnrID in context",
				Error:      dto.ErrorResponse{Message: "Organization ID not found", IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		sharpnrID, ok := sharpnrIDVal.(string)
		if !ok || sharpnrID == "" {
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid SharpnrID format",
				Error:      dto.ErrorResponse{Message: "Invalid SharpnrID in context", IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		// Fetch organization from DB
		var org models.Organization
		if err := database.DB.Where("sharpnr_id = ?", sharpnrID).First(&org).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
				StatusCode: http.StatusUnauthorized,
				Message:    "Organization not found",
				Error:      dto.ErrorResponse{Message: err.Error(), IsError: true},
				Timestamp:  time.Now(),
			})
			ctx.Abort()
			return
		}

		// Set the Org UUID in context
		ctx.Set("OrgUUID", org.ID)

		ctx.Next()
	}
}
