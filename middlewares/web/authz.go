package webMiddleware

import (
	"net/http"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequirePermission is the route-level authorization gate. It is the thin
// counterpart to calling s.sentinelEngine.IsAuthorized(...) inside a service:
// use it for coarse checks that depend only on the action + resource type (and
// optionally an :id path param). For decisions that need resource attributes or
// request context, authorize inside the service where that data is loaded.
//
// It MUST run after WebAuthMiddleware, which populates the org/user identity
// keys on the gin context. The engine fails closed, so any infrastructure error
// or an explicit deny results in 403.
//
// resourceIDParam is the name of the gin path parameter holding the resource id
// (e.g. "id"); pass "" for org-wide actions, in which case the org id is used
// as the resource handle.
func RequirePermission(
	engine *sentinelengine.Sentinel,
	action sentinel_actions.Name,
	resourceType string,
	resourceIDParam string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orgID, userID, ok := identityFromContext(ctx)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
			})
			return
		}

		resourceID := orgID.String()
		if resourceIDParam != "" {
			if v := ctx.Param(resourceIDParam); v != "" {
				resourceID = v
			}
		}

		decision, err := engine.IsAuthorized(ctx.Request.Context(), sentinelengine.AuthorizeInput{
			OrgID:  orgID,
			UserID: userID,
			Action: string(action),
			Resource: sentinelengine.Resource{
				Type: resourceType,
				ID:   resourceID,
			},
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}
		if !decision.Allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		ctx.Next()
	}
}

// identityFromContext pulls the org and user ids that WebAuthMiddleware stored.
// Both were set as uuid.UUID (claims.OrgID / claims.UserID).
func identityFromContext(ctx *gin.Context) (orgID, userID uuid.UUID, ok bool) {
	rawOrg, exists := ctx.Get(OrgUUIDContextKey)
	if !exists {
		return uuid.Nil, uuid.Nil, false
	}
	rawUser, exists := ctx.Get(UserIDContextKey)
	if !exists {
		return uuid.Nil, uuid.Nil, false
	}

	orgID, ok1 := rawOrg.(uuid.UUID)
	userID, ok2 := rawUser.(uuid.UUID)
	if !ok1 || !ok2 {
		return uuid.Nil, uuid.Nil, false
	}
	return orgID, userID, true
}
