package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/sharpnr-inc/sharpnr-engine/utils/types"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(empID, orgID uint) (string, error) {
	claims := &types.Claims{
		EmployeeID:     empID,
		OrganizationID: orgID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorisation")
		if !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer")
		claims := &types.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("employee_id", claims.EmployeeID)
		c.Set("organization_id", claims.OrganizationID)
		c.Next()

	}
}
