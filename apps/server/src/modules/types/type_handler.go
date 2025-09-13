package types

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type typeHandler struct {
	handler.GenericHandler[Type, TypeRequestDTO, TypeResponseDTO]
}

func NewTypeHandler(service typeService) *typeHandler {
	return &typeHandler{
		handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateTypeRequestDTO(c *gin.Context, dto *TypeRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[TypeResponseDTO]

// CreateType godoc
// @Summary Create a new Type
// @Description Create a new Type with the provided details
// @Tags types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param type body TypeRequestDTO true "Type creation details"
// @Success 201 {object} response.GenericResponse[TypeResponseDTO] "Type created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /types [post]
func (h *typeHandler) CreateType(c *gin.Context) {
	h.CreateHandler(c, ValidateTypeRequestDTO)
}

// GetTypeByID godoc
// @Summary Get a Type by ID
// @Description Get Type details by their unique ID
// @Tags types
// @Produce json
// @Security BearerAuth
// @Param id path string true "Type ID"
// @Success 200 {object} response.GenericResponse[TypeResponseDTO] "Type found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Type ID"
// @Failure 404 {object} response.GenericResponse[any] "Type not found"
// @Router /types/{id} [get]
func (h *typeHandler) GetTypeByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllTypes godoc
// @Summary Get all Types
// @Description Get a list of all registered Types
// @Tags types
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]TypeResponseDTO] "List of Types"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /types [get]
func (h *typeHandler) GetAllTypes(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateType godoc
// @Summary Update an existing Type
// @Description Update Type details by their unique ID
// @Tags types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Type ID"
// @Param type body TypeRequestDTO true "Type update details"
// @Success 200 {object} response.GenericResponse[TypeResponseDTO] "Type updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Type ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /types/{id} [put]
func (h *typeHandler) UpdateType(c *gin.Context) {
	h.UpdateHandler(c, ValidateTypeRequestDTO)
}

// DeleteType godoc
// @Summary Delete a Type
// @Description Delete a Type by their unique ID
// @Tags types
// @Produce json
// @Security BearerAuth
// @Param id path string true "Type ID"
// @Success 204 {object} response.GenericResponse[any] "Type deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Type ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /types/{id} [delete]
func (h *typeHandler) DeleteType(c *gin.Context) {
	h.DeleteHandler(c)
}
