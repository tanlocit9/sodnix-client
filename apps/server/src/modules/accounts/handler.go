package accounts

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	handler.GenericHandler[Account, AccountRequestDTO, AccountResponseDTO]
}

func NewAccountHandler(service accountService) *accountHandler {
	return &accountHandler{
		GenericHandler: handler.NewGenericHandler(service.GenericService),
	}
}

func ValidateAccountRequestDTO(c *gin.Context, dto *AccountRequestDTO) string {
	if err := c.ShouldBindJSON(dto); err != nil {
		return err.Error()
	}
	return ""
}

var _ response.GenericResponse[AccountResponseDTO]

// CreateAccount godoc
// @Summary Create a new Account
// @Description Create a new Account with the provided details
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param account body AccountRequestDTO true "Account creation details"
// @Success 201 {object} response.GenericResponse[AccountResponseDTO] "Account created successfully"
// @Failure 400 {object} response.GenericResponse[any] "Bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /accounts [post]
func (h *accountHandler) CreateAccount(c *gin.Context) {
	h.CreateHandler(c, ValidateAccountRequestDTO)
}

// GetAccountByID godoc
// @Summary Get an Account by ID
// @Description Get Account details by their unique ID
// @Tags accounts
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 200 {object} response.GenericResponse[AccountResponseDTO] "Account found"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Account ID"
// @Failure 404 {object} response.GenericResponse[any] "Account not found"
// @Router /accounts/{id} [get]
func (h *accountHandler) GetAccountByID(c *gin.Context) {
	h.GetByIDHandler(c)
}

// GetAllAccounts godoc
// @Summary Get all Accounts
// @Description Get a list of all registered Accounts
// @Tags accounts
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success 200 {object} response.GenericResponse[[]AccountResponseDTO] "List of Accounts"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /accounts [get]
func (h *accountHandler) GetAllAccounts(c *gin.Context) {
	h.GetAllHandler(c)
}

// UpdateAccount godoc
// @Summary Update an existing Account
// @Description Update Account details by their unique ID
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Param account body AccountRequestDTO true "Account update details"
// @Success 200 {object} response.GenericResponse[AccountResponseDTO] "Account updated successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Account ID or bad request"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /accounts/{id} [put]
func (h *accountHandler) UpdateAccount(c *gin.Context) {
	h.UpdateHandler(c, ValidateAccountRequestDTO)
}

// DeleteAccount godoc
// @Summary Delete an Account
// @Description Delete an Account by their unique ID
// @Tags accounts
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 204 {object} response.GenericResponse[any] "Account deleted successfully"
// @Failure 400 {object} response.GenericResponse[any] "Invalid Account ID"
// @Failure 500 {object} response.GenericResponse[any] "Internal server error"
// @Router /accounts/{id} [delete]
func (h *accountHandler) DeleteAccount(c *gin.Context) {
	h.DeleteHandler(c)
}
