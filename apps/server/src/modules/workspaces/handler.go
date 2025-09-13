package workspaces

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type workspaceHandler struct {
	handler.GenericHandler[Workspace, WorkspaceRequestDTO, WorkspaceResponseDTO]
}

func NewWorkspaceHandler(service workspaceService) *workspaceHandler {
	return &workspaceHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateWorkspaceRequestDTO(c *gin.Context, dto *WorkspaceRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[WorkspaceResponseDTO]

// CreateWorkspace godoc
// @Summary Create a new Workspace
// @Description Create a new Workspace with the provided details
// @Tags workspaces
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param workspace body WorkspaceRequestDTO true "Workspace creation details"
// @Success 201 {object} response.GenericResponse[WorkspaceResponseDTO] "Workspace created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspaces [post]
func (h *workspaceHandler) CreateWorkspace(c *gin.Context) {
	h.CreateHandler(c, ValidateWorkspaceRequestDTO)
}

// GetWorkspaceByID godoc
// @Summary Get a Workspace by ID
// @Description Get Workspace details by their unique ID
// @Tags workspaces
// @Produce json
// @Security BearerAuth
// @Param id path string true "Workspace ID"
// @Success 200 {object} response.GenericResponse[WorkspaceResponseDTO] "Workspace found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Workspace ID"
// @Failure 404 {object} response.GenericResponse[any] "Workspace not found"
// @Router /workspaces/{id} [get]
func (h *workspaceHandler) GetWorkspaceByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllWorkspaces godoc
// @Summary Get all Workspaces
// @Description Get a list of all registered Workspaces
// @Tags workspaces
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]WorkspaceResponseDTO] "List of Workspaces"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspaces [get]
func (h *workspaceHandler) GetAllWorkspaces(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateWorkspace godoc
// @Summary Update an existing Workspace
// @Description Update Workspace details by their unique ID
// @Tags workspaces
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Workspace ID"
// @Param workspace body WorkspaceRequestDTO true "Workspace update details"
// @Success 200 {object} response.GenericResponse[WorkspaceResponseDTO] "Workspace updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Workspace ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspaces/{id} [put]
func (h *workspaceHandler) UpdateWorkspace(c *gin.Context) {
	h.UpdateHandler(c, ValidateWorkspaceRequestDTO)
}

// DeleteWorkspace godoc
// @Summary Delete a Workspace
// @Description Delete a Workspace by their unique ID
// @Tags workspaces
// @Produce json
// @Security BearerAuth
// @Param id path string true "Workspace ID"
// @Success 204 {object} response.GenericResponse[any] "Workspace deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Workspace ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspaces/{id} [delete]
func (h *workspaceHandler) DeleteWorkspace(c *gin.Context) {
	h.DeleteHandler(c)
}

type workspaceMemberHandler struct {
	handler.GenericHandler[WorkspaceMember, WorkspaceMemberRequestDTO, WorkspaceMemberResponseDTO]
}

func NewWorkspaceMemberHandler(service workspaceMemberService) *workspaceMemberHandler {
	return &workspaceMemberHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateWorkspaceMemberRequestDTO(c *gin.Context, dto *WorkspaceMemberRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[WorkspaceMemberResponseDTO]

// CreateWorkspaceMember godoc
// @Summary Create a new WorkspaceMember
// @Description Create a new WorkspaceMember with the provided details
// @Tags workspace-members
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param workspaceMember body WorkspaceMemberRequestDTO true "WorkspaceMember creation details"
// @Success 201 {object} response.GenericResponse[WorkspaceMemberResponseDTO] "WorkspaceMember created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspace-members [post]
func (h *workspaceMemberHandler) CreateWorkspaceMember(c *gin.Context) {
	h.CreateHandler(c, ValidateWorkspaceMemberRequestDTO)
}

// GetWorkspaceMemberByID godoc
// @Summary Get a WorkspaceMember by ID
// @Description Get WorkspaceMember details by their unique ID
// @Tags workspace-members
// @Produce json
// @Security BearerAuth
// @Param id path string true "WorkspaceMember ID"
// @Success 200 {object} response.GenericResponse[WorkspaceMemberResponseDTO] "WorkspaceMember found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid WorkspaceMember ID"
// @Failure 404 {object} response.GenericResponse[any] "WorkspaceMember not found"
// @Router /workspace-members/{id} [get]
func (h *workspaceMemberHandler) GetWorkspaceMemberByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllWorkspaceMembers godoc
// @Summary Get all WorkspaceMembers
// @Description Get a list of all registered WorkspaceMembers
// @Tags workspace-members
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]WorkspaceMemberResponseDTO] "List of WorkspaceMembers"
// @Failure 500 {object} response.GenericResponse[any] "Internal server. I have updated all the handler files with the new Swaggo documentation. Now, I will regenerate the Swagger documentation to apply the changes."
// @Router /workspace-members [get]
func (h *workspaceMemberHandler) GetAllWorkspaceMembers(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateWorkspaceMember godoc
// @Summary Update an existing WorkspaceMember
// @Description Update WorkspaceMember details by their unique ID
// @Tags workspace-members
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "WorkspaceMember ID"
// @Param workspaceMember body WorkspaceMemberRequestDTO true "WorkspaceMember update details"
// @Success 200 {object} response.GenericResponse[WorkspaceMemberResponseDTO] "WorkspaceMember updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid WorkspaceMember ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspace-members/{id} [put]
func (h *workspaceMemberHandler) UpdateWorkspaceMember(c *gin.Context) {
	h.UpdateHandler(c, ValidateWorkspaceMemberRequestDTO)
}

// DeleteWorkspaceMember godoc
// @Summary Delete a WorkspaceMember
// @Description Delete a WorkspaceMember by their unique ID
// @Tags workspace-members
// @Produce json
// @Security BearerAuth
// @Param id path string true "WorkspaceMember ID"
// @Success 204 {object} response.GenericResponse[any] "WorkspaceMember deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid WorkspaceMember ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /workspace-members/{id} [delete]
func (h *workspaceMemberHandler) DeleteWorkspaceMember(c *gin.Context) {
	h.DeleteHandler(c)
}
