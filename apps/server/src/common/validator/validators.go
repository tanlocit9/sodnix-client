package validator

import (
	"github.com/gin-gonic/gin"
)

func ValidateRequestBody[T any](c *gin.Context, requestBody T) string {
	if err := c.ShouldBindJSON(requestBody); err != nil {
		return err.Error()
	}

	return ""
}
