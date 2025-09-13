package routes

import (
	"log"
	"sodnix/apps/server/src/modules/workspaces"

	"sodnix/apps/server/src/common/constants"
	"sodnix/apps/server/src/docs"
	"sodnix/apps/server/src/modules/accounts"
	"sodnix/apps/server/src/modules/auth"
	"sodnix/apps/server/src/modules/categories"
	"sodnix/apps/server/src/modules/transactions"
	"sodnix/apps/server/src/modules/types"
	"sodnix/apps/server/src/modules/users"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RegisterSwaggerRoute(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Norastro workflow API"
	docs.SwaggerInfo.Description = "This will contains all defined API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET(constants.BASE_API_GROUP+constants.API_SWAGGER_ACCEPT, ginSwagger.WrapHandler(swaggerFiles.Handler))

	swaggerURL := docs.SwaggerInfo.Host + docs.SwaggerInfo.BasePath + constants.API_SWAGGER_URL
	log.Println("Swagger page is running at " + swaggerURL)
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group(constants.BASE_API_GROUP)

	// Public routes
	auth.RegisterRoutes(api) // contains login/register

	// Protected routes
	protected := api.Group("/")
	// Will turn on later if needed for server-side auth
	// protected.Use(auth.SessionMiddleware())
	protected.Use(auth.JWTMiddleware())

	transactions.RegisterRoutes(protected, transactions.NewTransactionHandlers(db))

	types.RegisterRoutes(protected, types.NewTypeHandlers(db), types.NewTypeGroupHandlers(db))

	categories.RegisterRoutes(protected, categories.NewCategoryHandlers(db))

	users.RegisterRoutes(protected, users.NewUserHandlers(db))

	accounts.RegisterRoutes(protected, accounts.NewAccountHandlers(db))

	workspaces.RegisterRoutes(protected, workspaces.NewWorkspaceHandlers(db), workspaces.NewWorkspaceMemberHandlers(db))

	// Swagger
	RegisterSwaggerRoute(r)
}
