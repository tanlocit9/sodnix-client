package categories

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCategoryHandlers(db *gorm.DB) *categoryHandler {
	repository := NewCategoryRepository(db)
	categoryService := NewCategoryService(repository, NewMapper())
	return NewCategoryHandler(categoryService)
}

func RegisterRoutes(router *gin.RouterGroup, handlers *categoryHandler) {
	routes := router.Group(constants.API_PATH_CATEGORY)
	{
		routes.GET("/:id", handlers.GetCategoryByID)
		routes.GET("/", handlers.GetAllCategories)
		routes.POST("/", handlers.CreateCategory)
		routes.PUT("/:id", handlers.UpdateCategory)
		routes.DELETE("/:id", handlers.DeleteCategory)
	}
}
