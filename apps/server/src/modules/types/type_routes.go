package types

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTypeHandlers(db *gorm.DB) *typeHandler {
	typeRepo := NewTypeRepository(db)
	typeSvc := NewTypeService(*typeRepo, NewTypeMapper())

	return NewTypeHandler(typeSvc)
}

func RegisterTypeRoutes(rg *gin.RouterGroup, handlers *typeHandler) {
	types := rg.Group(constants.API_PATH_TYPE)
	{
		types.GET(":id", handlers.GetTypeByID)
		types.GET("", handlers.GetAllTypes)
		types.POST("", handlers.CreateType)
		types.PUT(":id", handlers.UpdateType)
		types.DELETE(":id", handlers.DeleteType)
	}
}
