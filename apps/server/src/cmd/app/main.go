package main

import (
	"sodnix/apps/server/src/common/converter"
	"sodnix/apps/server/src/config"
	"sodnix/apps/server/src/database"
	"sodnix/apps/server/src/modules/accounts"
	"sodnix/apps/server/src/modules/categories"
	"sodnix/apps/server/src/modules/transactions"
	"sodnix/apps/server/src/modules/types"
	"sodnix/apps/server/src/modules/users"
	"sodnix/apps/server/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// @title Norastro API
// @version 1.0
// @description src API for the Norastro Expense Management App.
// @termsOfService http://swagger.io/terms/

// @contact.name Norastro Dev Team
// @contact.url https://github.com/your-org/norastro
// @contact.email support@norastro.app

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()

	db := database.ConnectDatabase()

	database.Migrate(converter.ConcatToAny(categories.Models(), types.Models(), transactions.Models(), accounts.Models(), users.Models())...)

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routes.RegisterRoutes(r, db)

	r.Run(config.SERVER_HOST + ":" + config.SERVER_PORT)
}
