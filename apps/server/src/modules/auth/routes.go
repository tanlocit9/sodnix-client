package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", Login)
	authGroup.POST("/logout", Logout)
	authGroup.GET("/profile", JWTMiddleware(), Profile)
}
