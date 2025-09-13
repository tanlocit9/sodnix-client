package categories

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	handler.GenericHandler[Category, CategoryRequestDTO, CategoryResponseDTO]
}

func NewCategoryHandler(service categoryService) *categoryHandler {
	return &categoryHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateCategoryRequestDTO(c *gin.Context, dto *CategoryRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[CategoryResponseDTO]

// CreateCategory godoc
// @Summary Create a new Category
// @Description Create a new Category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category body CategoryRequestDTO true "Category creation details"
// @Success 201 {object} response.GenericResponse[CategoryResponseDTO] "Category created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /categories [post]
func (h *categoryHandler) CreateCategory(c *gin.Context) {
	h.CreateHandler(c, ValidateCategoryRequestDTO)
}

// GetCategoryByID godoc
// @Summary Get a Category by ID
// @Description Get Category details by their unique ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {object} response.GenericResponse[CategoryResponseDTO] "Category found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Category ID"
// @Failure 404 {object} response.GenericResponse[any] "Category not found"
// @Router /categories/{id} [get]
func (h *categoryHandler) GetCategoryByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllCategories godoc
// @Summary Get all Categories
// @Description Get a list of all registered Categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]CategoryResponseDTO] "List of Categories"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /categories [get]
func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateCategory godoc
// @Summary Update an existing Category
// @Description Update Category details by their unique ID
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param category body CategoryRequestDTO true "Category update details"
// @Success 200 {object} response.GenericResponse[CategoryResponseDTO] "Category updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Category ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /categories/{id} [put]
func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	h.UpdateHandler(c, ValidateCategoryRequestDTO)
}

// DeleteCategory godoc
// @Summary Delete a Category
// @Description Delete a Category by their unique ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 204 {object} response.GenericResponse[any] "Category deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Category ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /categories/{id} [delete]
func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	h.DeleteHandler(c)
}
