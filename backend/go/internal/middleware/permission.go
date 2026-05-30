package middleware

import (
	"net/http"
	"strings"

	"watch/backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func RequirePermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		authClaims, ok := claims.(*auth.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}
		if authClaims.Permissions == "all" {
			c.Next()
			return
		}
		for _, perm := range strings.Split(authClaims.Permissions, ",") {
			if strings.TrimSpace(perm) == requiredPermission {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}

func RequireAnyPermission(requiredPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		authClaims, ok := claims.(*auth.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}
		if authClaims.Permissions == "all" {
			c.Next()
			return
		}
		userPerms := strings.Split(authClaims.Permissions, ",")
		for _, userPerm := range userPerms {
			userPerm = strings.TrimSpace(userPerm)
			for _, required := range requiredPermissions {
				if userPerm == required {
					c.Next()
					return
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}
