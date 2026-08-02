package authenticationEngineControllers

import (
	"net/http"
	"time"

	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	"sharpnr-engine/pkg/cookieManager"
	sentinelContexts "sharpnr-engine/sentinel_contexts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *authenticationControllers) WebLogin(ctx *gin.Context) {

	var req authenticationEngineDtos.LoginRequest_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	// Login sits ahead of WebAuthMiddleware (no session yet), so stamp the
	// request IP here — LoginUser reads it back via sentinelContexts.Extract
	// to attribute the login's audit log entry.
	reqCtx := sentinelContexts.Inject(ctx.Request.Context(), sentinelContexts.AuthzContext{
		Request: sentinelContexts.RequestContext{
			Channel:     sentinelContexts.ChannelWeb,
			IPAddress:   ctx.ClientIP(),
			RequestedAt: time.Now(),
		},
	})

	res, err := s.authSvc.LoginUser(reqCtx, req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	authToken, err := s.jwtMgr.IssueWebAuthCookieToken(res.EmployeeID, res.OrgId)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
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
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	// Login sits ahead of mobile auth middleware (no session yet), so stamp
	// the request IP here — LoginUser reads it back via
	// sentinelContexts.Extract to attribute the login's audit log entry.
	reqCtx := sentinelContexts.Inject(ctx.Request.Context(), sentinelContexts.AuthzContext{
		Request: sentinelContexts.RequestContext{
			Channel:     sentinelContexts.ChannelMobile,
			IPAddress:   ctx.ClientIP(),
			RequestedAt: time.Now(),
		},
	})

	res, err := s.authSvc.LoginUser(reqCtx, req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	// Access is issued first — if it fails we haven't wasted the refresh issuance.
	// res.Role, not res.Org.Name: the role claim previously carried the
	// organisation name, which made the claim useless for authorisation.
	accessToken, err := s.jwtMgr.IssueAccess(res.EmployeeID, res.OrgId, res.Role)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	// The session row must exist before the token that references it is
	// handed out, otherwise the first refresh finds nothing to rotate.
	sessionID, jti, err := s.authSvc.StartMobileSession(reqCtx, res.EmployeeID, res.OrgId, res.Role)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := s.jwtMgr.IssueRefresh(res.EmployeeID, res.OrgId, sessionID, jti)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	res.AccessToken = &accessToken
	res.RefreshToken = &refreshToken

	utils.Success(ctx, http.StatusOK, "Logged in Successfully", res)
}
