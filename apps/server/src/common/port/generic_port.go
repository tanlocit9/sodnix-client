package port

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CrudServicePort[Model any, RequestDTO any, ResponseDTO any] interface {
	Create(c *gin.Context, input RequestDTO) (*ResponseDTO, error)
	GetByID(id uuid.UUID, opts *common.GetOptions) (*ResponseDTO, error)
	GetAll(params *common.ListQueryParams) (*response.GenericResponse[[]ResponseDTO], error)
	Update(c *gin.Context, id uuid.UUID, input RequestDTO) (*ResponseDTO, error)
	Delete(id uuid.UUID) error
}

type CrudRepositoryPort[T any] interface {
	Save(model *T) error
	FindByID(id uuid.UUID, opts *common.GetOptions) (*T, error)
	FindAll(params *common.ListQueryParams) ([]T, int64, error)
	Update(model *T) error
	Delete(id uuid.UUID) error
}
