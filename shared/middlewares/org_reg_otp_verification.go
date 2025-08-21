package middlewares

import (
	"net/http"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// OrgRegOtpVerificationMiddleware is a Gin middleware used to intercept and validate
// the "Sharpnr-Org-Otp-Auth-token" header during organization registration.
//
// Workflow:
//  1. Checks if the request contains the custom header "Sharpnr-Org-Otp-Auth-token".
//  2. Ensures the header value is in the format: "Bearer <token>".
//  3. Extracts the token part and validates it using utils.ValidateJWT().
//  4. If token is valid, extracts the "email" claim and stores it in Gin's context
//     (via c.Set("email", ...)) so that downstream handlers can access it.
//  5. If token is missing, invalid, or expired, the request is aborted with 401 Unauthorized.
//
// Example usage:
//
//	r := gin.Default()
//	r.POST("/verify-otp", OrgRegOtpVerificationMiddleware(), verifyOtpHandler)
//
// Inside verifyOtpHandler, you can retrieve the email like this:
//
//	email := c.GetString("email")
//
// Returns a gin.HandlerFunc which can be used with routes or groups.
func OrgRegOtpVerificationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract custom header
		otpHeader := c.GetHeader("Authorization")
		if otpHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Auth Token"})
			return
		}

		// Expect format: "Bearer <token>"
		parts := strings.Split(otpHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// Validate the JWT
		tokenString := parts[1]
		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Extract claims (assuming standard MapClaims)
		claims := token.Claims.(jwt.MapClaims)

		// Store email in Gin context for later handlers
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email not found in token"})
			return
		}

		// Continue to the next middleware/handler
		c.Next()
	}
}
