package transactions

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTransactionHandlers(db *gorm.DB) *transactionHandler {
	transactionRepository := NewTransactionRepository(db)
	transactionService := NewTransactionService(transactionRepository, NewMapper())
	return NewTransactionHandler(transactionService)
}

func RegisterRoutes(rg *gin.RouterGroup, handler *transactionHandler) {

	transactions := rg.Group(constants.API_PATH_TRANSACTION)
	transactions.GET(":id", handler.GetTransactionByIdHandler)
	transactions.GET("", handler.GetAllTransactionHandler)
	transactions.POST("", handler.CreateTransactionHandler)
	transactions.PUT(":id", handler.UpdateTransactionHandler)
	transactions.DELETE(":id", handler.DeleteTransactionHandler)
}
