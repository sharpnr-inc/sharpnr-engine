package coreutilitiesapi

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBusinessRegistrationType(c *gin.Context) {
	businessRegType := []constants.BusinessRegistrationType{
		constants.SoleProprietorship,
		constants.Partnership,
		constants.LimitedLiabilityPartnership,
		constants.PrivateLimitedCompany,
		constants.PublicLimitedCompany,
		constants.OnePersonCompany,
		constants.Section8Company,
		constants.CooperativeSociety,
		constants.Trust,
		constants.Society,
	}

	// ✅ Wrap inside a JSON object key
	wrappedData := map[string]any{
		"BUSINESS_REGISTRATION_TYPE": businessRegType,
	}

	// ✅ Build standardized response
	res := dto.Response[map[string]any]{
		StatusCode: http.StatusOK,
		Message:    "List of All Business Registration Types",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &wrappedData, // pointer to map
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, res)
}
