package transactions

import (
	"sodnix/apps/server/src/common/handler"
	"sodnix/apps/server/src/common/response"
	"sodnix/apps/server/src/common/validator"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	handler.GenericHandler[Transaction, TransactionRequestDTO, TransactionResponseDTO]
}

func NewTransactionHandler(service transactionService) *transactionHandler {
	return &transactionHandler{
		handler.NewGenericHandler(service.GenericService),
	}
}

// GetTransactionById godoc
// @Summary Get Transaction by ID
// @Tags transactions
// @Security BearerAuth
// @Param 		 id path string true "Transaction ID"
// @Success 	 200 {object} response.GetDataSuccess[TransactionResponseDTO]
// @Failure      400 {object} response.BadRequestError
// @Failure      404 {object} response.NotFoundError
// @Failure      500 {object} response.InternalServerError
// @Router /transactions/{id} [get]
func (h *transactionHandler) GetTransactionByIdHandler(c *gin.Context) {
	var _ response.GetDataSuccess[TransactionResponseDTO]
	h.GetByIDHandler(c)
}

// GetAllTransaction godoc
//
// @Summary	Get all transactions
// @Tags		transactions
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param sort_field query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order" default(desc)
// @Param search query string false "Search term"
// @Param preload query string false "Preload relations"
// @Success      200 {object} response.GetDataSuccess[[]TransactionResponseDTO]
// @Failure      400 {object} response.BadRequestError
// @Failure      404 {object} response.NotFoundError
// @Failure      500 {object} response.InternalServerError
// @Router		/transactions [get]
func (h *transactionHandler) GetAllTransactionHandler(c *gin.Context) {
	h.GetAllHandler(c)
}

// CreateTransaction godoc
// @Summary	Create an Transaction
// @Tags		transactions
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param      request body TransactionRequestDTO true "TransactionRequestDTO"
// @Success      201 {object} response.CreateSuccess[TransactionResponseDTO]
// @Failure      400 {object} response.BadRequestError
// @Failure      500 {object} response.InternalServerError
// @Router		/transactions [post]
func (h *transactionHandler) CreateTransactionHandler(c *gin.Context) {
	h.CreateHandler(c, func(c *gin.Context, req *TransactionRequestDTO) string {
		return validator.ValidateRequestBody(c, req)
	})
}

// UpdateTransaction godoc
// @Summary	Update an Transaction
// @Tags		transactions
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Transaction ID"
// @Param      request body TransactionRequestDTO true "TransactionRequestDTO"
// @Success      200 {object} response.CreateSuccess[TransactionResponseDTO]
// @Failure      400 {object} response.BadRequestError
// @Failure      404 {object} response.NotFoundError
// @Failure      500 {object} response.InternalServerError
// @Router		/transactions/{id} [put]
func (h *transactionHandler) UpdateTransactionHandler(c *gin.Context) {
	h.UpdateHandler(c, func(c *gin.Context, req *TransactionRequestDTO) string {
		return validator.ValidateRequestBody(c, req)
	})
}

// DeleteTransaction godoc
// @Summary Delete an Transaction
// @Tags transactions
// @Security BearerAuth
// @Param id path string true "Transaction ID"
// @Success 204 {object} response.GenericResponse[any] "Transaction deleted successfully"
// @Failure 400 {object} response.BadRequestError
// @Failure 500 {object} response.InternalServerError
// @Router /transactions/{id} [delete]
func (h *transactionHandler) DeleteTransactionHandler(c *gin.Context) {
	h.DeleteHandler(c)
}
