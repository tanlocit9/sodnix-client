package types

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, typeHandler *typeHandler, typeGroupHandler *typeGroupHandler) {
	RegisterTypeRoutes(rg, typeHandler)
	RegisterTypeGroupRoutes(rg, typeGroupHandler)
}
