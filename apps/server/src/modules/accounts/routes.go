package accounts

import (
	"sodnix/apps/server/src/common/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAccountHandlers(db *gorm.DB) *accountHandler {
	repository := NewAccountRepository(db)
	accountService := NewAccountService(repository, NewMapper())
	return NewAccountHandler(accountService)
}

func RegisterRoutes(router *gin.RouterGroup, handlers *accountHandler) {
	routes := router.Group(constants.API_PATH_ACCOUNT)
	{
		routes.POST("/", handlers.CreateAccount)
		routes.GET("/:id", handlers.GetAccountByID)
		routes.GET("/", handlers.GetAllAccounts)
		routes.PUT("/:id", handlers.UpdateAccount)
		routes.DELETE("/:id", handlers.DeleteAccount)
	}
}
