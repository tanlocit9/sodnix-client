package auth

import (
	"fmt"
	"net/http"
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userID, ok := GetSessionUserID(c); ok {
			c.Set(constants.AUTH_USER_ID_KEY, userID)
		} else {
			// If session user ID is not found, clear the session to prevent stale data
			ClearSession(c)
		}
		c.Next()
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip if already set by session
		if _, exists := c.Get(constants.AUTH_USER_ID_KEY); exists {
			c.Next()
			return
		}

		tokenStr := c.GetHeader(constants.AUTH_HEADER)
		if tokenStr != "" {
			if claims, err := ParseJWT(tokenStr); err == nil {
				// Ensure the UserID from claims is a valid UUID before setting it
				fmt.Println(claims)
				if claims.UserID != uuid.Nil {
					c.Set(constants.AUTH_USER_ID_KEY, claims.UserID)
					c.Set(constants.AUTH_EMAIL_KEY, claims.Email)
					c.Set(constants.AUTH_DISPLAY_NAME_KEY, claims.DisplayName)
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}
