package response

import "github.com/gin-gonic/gin"

type Paging struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type GenericResponse[T any] struct {
	Code    int     `json:"code" example:"200"`
	Message string  `json:"message" example:"OK"`
	Data    *T      `json:"data,omitempty"`
	Paging  *Paging `json:"paging,omitempty"`
}

func Response[T any](c *gin.Context, response GenericResponse[T]) {
	c.JSON(response.Code, response)
}
