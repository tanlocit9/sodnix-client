package workspaces

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewWorkspaceHandlers(db *gorm.DB) *workspaceHandler {
	repository := NewWorkspaceRepository(db)
	workspaceService := NewWorkspaceService(repository, NewWorkspaceMapper())
	return NewWorkspaceHandler(workspaceService)
}

func NewWorkspaceMemberHandlers(db *gorm.DB) *workspaceMemberHandler {
	repository := NewWorkspaceMemberRepository(db)
	workspaceMemberService := NewWorkspaceMemberService(repository, NewWorkspaceMemberMapper())
	return NewWorkspaceMemberHandler(workspaceMemberService)
}

func RegisterRoutes(router *gin.RouterGroup, workspaceHandlers *workspaceHandler, workspaceMemberHandlers *workspaceMemberHandler) {
	workspaceRoutes := router.Group(constants.API_PATH_WORKSPACE)
	{
		workspaceRoutes.POST("/", workspaceHandlers.CreateWorkspace)
		workspaceRoutes.GET("/:id", workspaceHandlers.GetWorkspaceByID)
		workspaceRoutes.GET("/", workspaceHandlers.GetAllWorkspaces)
		workspaceRoutes.PUT("/:id", workspaceHandlers.UpdateWorkspace)
		workspaceRoutes.DELETE("/:id", workspaceHandlers.DeleteWorkspace)
	}

	workspaceMemberRoutes := router.Group(constants.API_PATH_WORKSPACE_MEMBER)
	{
		workspaceMemberRoutes.POST("/", workspaceMemberHandlers.CreateWorkspaceMember)
		workspaceMemberRoutes.GET("/:id", workspaceMemberHandlers.GetWorkspaceMemberByID)
		workspaceMemberRoutes.GET("/", workspaceMemberHandlers.GetAllWorkspaceMembers)
		workspaceMemberRoutes.PUT("/:id", workspaceMemberHandlers.UpdateWorkspaceMember)
		workspaceMemberRoutes.DELETE("/:id", workspaceMemberHandlers.DeleteWorkspaceMember)
	}
}
