package hospitality_utilities

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetHospitalityDomainStencil(ctx *gin.Context) {

	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return
	}

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	// 3. Domain validation (case-insensitive)
	domain := ctx.Param("domain")
	if strings.ToUpper(domain) != "HOSPITALITY" {
		utils.Failure(
			ctx,
			http.StatusBadRequest,
			utils.CodeValidationFailed,
			errors.New("invalid domain: only HOSPITALITY is supported"),
		)
		return
	}

	// 4. Response payload
	response := gin.H{
		"domain":         "HOSPITALITY",
		"modules":        HospitalityModules,
		"infrastructure": HospitalityDomainInfraModules,
	}

	utils.Success(
		ctx,
		http.StatusOK,
		"Success",
		response,
	)

}
