package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ahmad20/bri-mini-project/repositories"
	"github.com/gin-gonic/gin"
)

type authInfo struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func AuthMiddleware(authService AuthInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := authService.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		authInfo := authInfo{
			Username: claims.Username,
			Role:     claims.Role,
		}
		c.Set("AuthInfo", authInfo)

		c.Next()
	}
}

func RoleAuthorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		info, exists := c.Get("AuthInfo")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		userName := info.(authInfo).Username
		userRole := info.(authInfo).Role
		if !IsRoleAllowed(userRole, allowedRoles) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("Username", userName)
		c.Next()
	}
}
func StatusAuthorization(repo repositories.AccountRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("Username")
		account, err := repo.SearchByUsername(fmt.Sprint(username))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			c.Abort()
			return
		}

		if account.ApprovalStatus == "waiting" || account.Status == "inactive" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "account not approved or inactive"})
			c.Abort()
			return
		}
		c.Next()
	}

}
func IsRoleAllowed(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
