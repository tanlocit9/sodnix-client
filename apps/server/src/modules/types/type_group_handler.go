package types

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type typeGroupHandler struct {
	handler.GenericHandler[TypeGroup, TypeGroupRequestDTO, TypeGroupResponseDTO]
}

func NewTypeGroupHandler(service typeGroupService) *typeGroupHandler {
	return &typeGroupHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateTypeGroupRequestDTO(c *gin.Context, dto *TypeGroupRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[TypeGroupResponseDTO]

// CreateTypeGroup godoc
// @Summary Create a new TypeGroup
// @Description Create a new TypeGroup with the provided details
// @Tags type-groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param typeGroup body TypeGroupRequestDTO true "TypeGroup creation details"
// @Success 201 {object} response.GenericResponse[TypeGroupResponseDTO] "TypeGroup created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /type-groups [post]
func (h *typeGroupHandler) CreateTypeGroup(c *gin.Context) {
	h.CreateHandler(c, ValidateTypeGroupRequestDTO)
}

// GetTypeGroupByID godoc
// @Summary Get a TypeGroup by ID
// @Description Get TypeGroup details by their unique ID
// @Tags type-groups
// @Produce json
// @Security BearerAuth
// @Param id path string true "TypeGroup ID"
// @Success 200 {object} response.GenericResponse[TypeGroupResponseDTO] "TypeGroup found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid TypeGroup ID"
// @Failure 404 {object} response.GenericResponse[any] "TypeGroup not found"
// @Router /type-groups/{id} [get]
func (h *typeGroupHandler) GetTypeGroupByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllTypeGroups godoc
// @Summary Get all TypeGroups
// @Description Get a list of all registered TypeGroups
// @Tags type-groups
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]TypeGroupResponseDTO] "List of TypeGroups"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /type-groups [get]
func (h *typeGroupHandler) GetAllTypeGroups(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateTypeGroup godoc
// @Summary Update an existing TypeGroup
// @Description Update TypeGroup details by their unique ID
// @Tags type-groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "TypeGroup ID"
// @Param typeGroup body TypeGroupRequestDTO true "TypeGroup update details"
// @Success 200 {object} response.GenericResponse[TypeGroupResponseDTO] "TypeGroup updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid TypeGroup ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /type-groups/{id} [put]
func (h *typeGroupHandler) UpdateTypeGroup(c *gin.Context) {
	h.UpdateHandler(c, ValidateTypeGroupRequestDTO)
}

// DeleteTypeGroup godoc
// @Summary Delete a TypeGroup
// @Description Delete a TypeGroup by their unique ID
// @Tags type-groups
// @Produce json
// @Security BearerAuth
// @Param id path string true "TypeGroup ID"
// @Success 204 {object} response.GenericResponse[any] "TypeGroup deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid TypeGroup ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /type-groups/{id} [delete]
func (h *typeGroupHandler) DeleteTypeGroup(c *gin.Context) {
	h.DeleteHandler(c)
}
