package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RoleAuthorization(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract the token from the header (format: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // means no Bearer prefix was found
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format, expected 'Bearer <token>'"})
			c.Abort()
			return
		}

		// Parse the JWT token (without verification for example purposes)
		// In production, you MUST verify the token with your secret key
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Get the role from claims
		// The claim path might differ based on your JWT structure
		roleClaim := claims["http://schemas.microsoft.com/ws/2008/06/identity/claims/role"]
		userRole, ok := roleClaim.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role claim not found or invalid"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied - insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
