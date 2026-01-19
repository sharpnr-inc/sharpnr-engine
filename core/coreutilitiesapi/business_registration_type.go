package coreutilitiesapi

import (
	"net/http"

	"sharpnr-engine/core/constants"
	"sharpnr-engine/shared/utils"

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

	// Wrap inside a JSON object key (if frontend expects this shape)
	wrappedData := map[string]any{
		"BUSINESS_REGISTRATION_TYPE": businessRegType,
	}

	utils.Success(
		c,
		http.StatusOK,
		"List of All Business Registration Types",
		wrappedData,
	)
}
