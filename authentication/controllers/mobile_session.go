package authenticationEngineControllers

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	authenticationEngineServices "sharpnr-engine/authentication/services"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/gin-gonic/gin"
)

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// MobileRefresh exchanges a refresh token for a new access/refresh pair.
//
//	POST /api/v1/mobile/identity/auth/refresh
//
// The refresh token is rotated on every call: the token supplied here is
// retired the moment it is redeemed and a new one is returned alongside the
// access token. Clients must persist both from every response — replaying a
// previous refresh token is treated as theft and signs the user out of every
// device (see RotateMobileSession).
//
// Sits outside MobileMiddleware: the caller's access token is expired by
// definition, so requiring one would deadlock the flow.
func (s *authenticationControllers) MobileRefresh(ctx *gin.Context) {

	var body refreshTokenRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "bad_request",
			"message": "refresh_token is required",
		})
		return
	}

	// This endpoint runs ahead of the auth middleware, so nothing has stamped
	// the request IP yet. The audit entry written on replay detection needs it.
	reqCtx := sentinelContexts.Inject(ctx.Request.Context(), sentinelContexts.AuthzContext{
		Request: sentinelContexts.RequestContext{
			Channel:     sentinelContexts.ChannelMobile,
			IPAddress:   ctx.ClientIP(),
			RequestedAt: time.Now(),
		},
	})

	// 1️⃣ Verify the refresh token's signature, issuer, audience and expiry.
	claims, err := s.jwtMgr.VerifyRefresh(body.RefreshToken)
	if err != nil {
		if errors.Is(err, jwt_manager.ErrExpiredToken) {
			slog.Info("mobile refresh token expired",
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "session_expired",
				"message": "Session expired, please log in again",
			})
			return
		}

		slog.Warn("mobile refresh failed",
			"reason", "invalid_refresh_token",
			"error", err.Error(),
			"path", ctx.Request.URL.Path,
			"ip", ctx.ClientIP(),
		)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "unauthorized",
			"message": "Invalid refresh token",
		})
		return
	}

	// 2️⃣ A signed token with no jti predates rotation and has no session row.
	presentedJTI, err := claims.JTI()
	if err != nil {
		slog.Warn("mobile refresh failed",
			"reason", "missing_jti",
			"user_id", claims.UserID,
			"ip", ctx.ClientIP(),
		)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "session_expired",
			"message": "Session expired, please log in again",
		})
		return
	}

	// 3️⃣ Rotate. This is the authorisation step — the signature above only
	// proves the token was minted by us, not that it is still current.
	newJTI, role, err := s.authSvc.RotateMobileSession(
		reqCtx,
		claims.SessionID,
		presentedJTI,
		claims.UserID,
		claims.OrgID,
	)
	if err != nil {
		s.abortRotationFailure(ctx, err, claims.UserID)
		return
	}

	// 4️⃣ Mint the new pair. Role comes from the session row, so it survives
	// refresh instead of being dropped as it was previously.
	accessToken, err := s.jwtMgr.IssueAccess(claims.UserID, claims.OrgID, role)
	if err != nil {
		slog.Error("failed to issue access token",
			"user_id", claims.UserID,
			"error", err.Error(),
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "server_error",
			"message": "Failed to issue access token",
		})
		return
	}

	refreshToken, err := s.jwtMgr.IssueRefresh(claims.UserID, claims.OrgID, claims.SessionID, newJTI)
	if err != nil {
		// The session already advanced to newJTI, so the token the client
		// still holds is dead. Nothing to roll back to — surface the failure
		// and let the client re-login rather than hand back a stale pair.
		slog.Error("failed to issue refresh token after rotation",
			"user_id", claims.UserID,
			"session_id", claims.SessionID,
			"error", err.Error(),
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "server_error",
			"message": "Failed to issue refresh token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// abortRotationFailure maps a rotation error onto the response. Every case is
// a 401 — the distinction is in the code, which tells the client whether this
// was an ordinary expiry or something that warrants surfacing to the user.
func (s *authenticationControllers) abortRotationFailure(
	ctx *gin.Context,
	err error,
	userID any,
) {
	switch {
	case errors.Is(err, authenticationEngineServices.ErrRefreshTokenReuse):
		// Loud on purpose: this is the one case that indicates a stolen token.
		slog.Warn("mobile refresh token reuse detected — all sessions revoked",
			"user_id", userID,
			"path", ctx.Request.URL.Path,
			"ip", ctx.ClientIP(),
		)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "token_reuse_detected",
			"message": "Session security violation, please log in again",
		})

	case errors.Is(err, authenticationEngineServices.ErrMobileSessionRevoked):
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "session_revoked",
			"message": "Session has been revoked, please log in again",
		})

	case errors.Is(err, authenticationEngineServices.ErrMobileSessionExpired),
		errors.Is(err, authenticationEngineServices.ErrMobileSessionNotFound):
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "session_expired",
			"message": "Session expired, please log in again",
		})

	default:
		slog.Error("mobile refresh failed",
			"reason", "rotation_error",
			"user_id", userID,
			"error", err.Error(),
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "server_error",
			"message": "Failed to refresh session",
		})
	}
}

/* -------------------------------------------------------------------------- */
/*                            Mobile Logout Handler                           */
/* -------------------------------------------------------------------------- */

// MobileLogout revokes the session behind a refresh token.
//
//	POST /api/v1/mobile/identity/auth/logout
//
// Takes the refresh token rather than the access token so logout still works
// once the access token has expired — the common case for an app resuming
// from the background straight into a sign-out.
func (s *authenticationControllers) MobileLogout(ctx *gin.Context) {

	var body refreshTokenRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    "bad_request",
			"message": "refresh_token is required",
		})
		return
	}

	claims, err := s.jwtMgr.VerifyRefresh(body.RefreshToken)
	if err != nil {
		// An expired token means the session is already unusable, so the
		// caller's intent is satisfied. Reporting failure would only push
		// clients into retry loops over an already-dead session.
		if errors.Is(err, jwt_manager.ErrExpiredToken) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Logged out"})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "unauthorized",
			"message": "Invalid refresh token",
		})
		return
	}

	if err := s.authSvc.EndMobileSession(ctx.Request.Context(), claims.SessionID); err != nil {
		slog.Error("failed to revoke mobile session",
			"user_id", claims.UserID,
			"session_id", claims.SessionID,
			"error", err.Error(),
		)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "server_error",
			"message": "Failed to log out",
		})
		return
	}

	slog.Info("mobile session revoked",
		"user_id", claims.UserID,
		"session_id", claims.SessionID,
		"ip", ctx.ClientIP(),
	)

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
