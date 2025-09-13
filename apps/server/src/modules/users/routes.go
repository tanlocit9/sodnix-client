package users

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserHandlers(db *gorm.DB) *userHandler {
	repository := NewUserRepository(db)
	userService := NewUserService(repository, NewMapper())
	return NewHandlers(userService)
}

func RegisterRoutes(router *gin.RouterGroup, handlers *userHandler) {
	routes := router.Group(constants.API_PATH_USER)
	{
		routes.POST("/", handlers.CreateUser)
		routes.GET("/:id", handlers.GetUserByID)
		routes.GET("/", handlers.GetAllUsers)
		routes.PUT("/:id", handlers.UpdateUser)
		routes.DELETE("/:id", handlers.DeleteUser)
	}
}
