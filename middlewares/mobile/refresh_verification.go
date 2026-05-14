package mobile_middleware

import (
	"errors"
	"log/slog"
	"net/http"

	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

// RefreshHandler is the endpoint the mobile client calls when it receives
// a "token_expired" response from MobileMiddleware.
//
// POST /api/v1/mobile/auth/refresh
//
// Flow:
//  1. Client receives 401 with code "token_expired"
//  2. Client sends refresh token to this endpoint
//  3. Handler verifies the refresh token
//  4. Handler issues a new access token and returns it
//  5. Client retries the original request with the new access token
//
// The refresh token itself is long-lived (default 24h) and is NOT rotated
// here — add rotation if you want stronger security (invalidate old refresh
// token and issue a new one on every use).
func RefreshHandler(jwtMgr jwt_manager.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "bad_request",
				"message": "refresh_token is required",
			})
			return
		}

		// 1️⃣ Verify the refresh token
		claims, err := jwtMgr.VerifyRefresh(body.RefreshToken)
		if err != nil {
			if errors.Is(err, jwt_manager.ErrExpiredToken) {
				// Refresh token itself has expired — full re-login required.
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

		// 2️⃣ TODO: look up claims.SessionID in your sessions table.
		// If the session has been revoked (e.g. user logged out on another device),
		// return 401 here before issuing a new token.
		//
		// Example:
		//   session, err := sessionRepo.FindByID(ctx, claims.SessionID)
		//   if err != nil || session.RevokedAt != nil {
		//       ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//           "code":    "session_revoked",
		//           "message": "Session has been revoked",
		//       })
		//       return
		//   }

		// 3️⃣ Issue a fresh access token
		// Role is not stored in the refresh token — fetch it from DB or
		// use a default. Adjust this to match your auth model.
		newAccessToken, err := jwtMgr.IssueAccess(claims.UserID, claims.OrgID, "")
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

		ctx.JSON(http.StatusOK, gin.H{
			"access_token": newAccessToken,
		})
	}
}
