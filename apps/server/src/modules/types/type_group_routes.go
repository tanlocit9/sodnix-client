package types

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTypeGroupHandlers(db *gorm.DB) *typeGroupHandler {
	groupRepo := NewTypeGroupRepository(db)
	groupSvc := NewTypeGroupService(*groupRepo, NewTypeGroupMapper())

	return NewTypeGroupHandler(groupSvc)
}

func RegisterTypeGroupRoutes(rg *gin.RouterGroup, handlers *typeGroupHandler) {
	typeGroups := rg.Group(constants.API_PATH_TYPE_GROUP)
	{
		typeGroups.GET(":id", handlers.GetTypeGroupByID)
		typeGroups.GET("", handlers.GetAllTypeGroups)
		typeGroups.POST("", handlers.CreateTypeGroup)
		typeGroups.PUT(":id", handlers.UpdateTypeGroup)
		typeGroups.DELETE(":id", handlers.DeleteTypeGroup)
	}
}
