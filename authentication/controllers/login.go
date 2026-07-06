package authenticationEngineControllers

import (
	"net/http"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	"sharpnr-engine/pkg/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *authenticationControllers) WebLogin(ctx *gin.Context) {

	var req authenticationEngineDtos.LoginRequest_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Invalid Credentials", err)
		return
	}

	res, err := s.authSvc.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	authToken, err := s.jwtMgr.IssueWebAuthCookieToken(res.EmployeeID, res.OrgId)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	s.cookieMgr.Set(ctx, cookieManager.Auth, authToken)
	s.cookieMgr.Set(ctx, cookieManager.CSRF, uuid.New().String())

	utils.Success(ctx, http.StatusOK, "Login Successfully", res)
}

/* -------------------------------------------------------------------------- */
/*                           Mobile Login Controller                          */
/* -------------------------------------------------------------------------- */

func (s *authenticationControllers) MobileLogin(ctx *gin.Context) {
	var req authenticationEngineDtos.LoginRequest_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Invalid Credentials", err)
		return
	}

	res, err := s.authSvc.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	// Access is issued first — if it fails we haven't wasted the refresh issuance
	accessToken, err := s.jwtMgr.IssueAccess(res.EmployeeID, res.OrgId, res.Org.Name)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	refreshToken, err := s.jwtMgr.IssueRefresh(res.EmployeeID, res.OrgId)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	res.AccessToken = &accessToken
	res.RefreshToken = &refreshToken

	utils.Success(ctx, http.StatusOK, "Logged in Successfully", res)
}
