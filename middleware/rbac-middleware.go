package middlewares

import (
	"go/auth/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(allowedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get(constants.ContextRole)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no role"})
			return
		}
		role := roleVal.(string)
		if role != allowedRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
