package users

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	handler.GenericHandler[User, UserRequestDTO, UserResponseDTO]
}

func NewHandlers(service userService) *userHandler {
	return &userHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateUserRequestDTO(c *gin.Context, dto *UserRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[UserResponseDTO]

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequestDTO true "User creation details"
// @Success 201 {object} response.GenericResponse[UserResponseDTO] "User created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /users [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	h.CreateHandler(c, ValidateUserRequestDTO)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get user details by their unique ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} response.GenericResponse[UserResponseDTO] "User found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid user ID"
// @Failure 404 {object} response.GenericResponse[any] "User not found"
// @Router /users/{id} [get]
func (h *userHandler) GetUserByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all registered users
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]UserResponseDTO] "List of users"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /users [get]
func (h *userHandler) GetAllUsers(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update user details by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param user body UserRequestDTO true "User update details"
// @Success 200 {object} response.GenericResponse[UserResponseDTO] "User updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid user ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /users/{id} [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	h.UpdateHandler(c, ValidateUserRequestDTO)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their unique ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 204 {object} response.GenericResponse[any] "User deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid user ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /users/{id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	h.DeleteHandler(c)
}
