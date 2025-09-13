package handler

import (
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/common/response"
	"sodnix/apps/server/src/common/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenericHandler struct holds the CrudServicePort for dependency injection
type GenericHandler[Model any, RequestDTO any, ResponseDTO any] struct {
	service service.GenericService[Model, RequestDTO, ResponseDTO]
}

// NewGenericHandler creates a new instance of GenericHandler
func NewGenericHandler[Model any, RequestDTO any, ResponseDTO any](service service.GenericService[Model, RequestDTO, ResponseDTO]) GenericHandler[Model, RequestDTO, ResponseDTO] {
	return GenericHandler[Model, RequestDTO, ResponseDTO]{service}
}

// GetByIDHandler handles fetching a single resource by ID
func (h *GenericHandler[Model, RequestDTO, ResponseDTO]) GetByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.Response(c, response.BadRequestResponse[string]("Invalid ID format"))
		return
	}

	// Extract preloads from query parameters
	preloads := c.Query("preloads")

	opts := &common.GetOptions{
		Preload: preloads,
	}

	resp, err := h.service.GetByID(id, opts)
	if err != nil {
		response.Response(c, response.NotFoundResponse[string](err.Error()))
		return
	}

	response.Response(c, response.GetDataSuccessResponse(resp))
}

// GetAllHandler handles fetching all resources
func (h *GenericHandler[Model, RequestDTO, ResponseDTO]) GetAllHandler(c *gin.Context) {
	var params common.ListQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Response(c, response.BadRequestResponse[string](err.Error()))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.SortField == "" {
		params.SortField = "created_at"
	}
	if params.SortOrder == "" {
		params.SortOrder = "desc"
	}

	resps, err := h.service.GetAll(&params)
	if err != nil {
		response.Response(c, response.InternalServerErrorResponse[string](err.Error()))
		return
	}

	response.Response(c, *resps)
}

// CreateHandler handles creating a new resource
func (h *GenericHandler[Model, RequestDTO, ResponseDTO]) CreateHandler(
	c *gin.Context,
	validateFunc func(*gin.Context, *RequestDTO) string,
) {
	var req RequestDTO
	if err := validateFunc(c, &req); err != "" {
		response.Response(c, response.BadRequestResponse[Model](err))
		return
	}

	resp, err := h.service.Create(c, req)
	if err != nil {
		response.Response(c, response.BadRequestResponse[string](err.Error()))
		return
	}

	response.Response(c, response.CreateDataSuccessResponse(resp))
}

// UpdateHandler handles updating an existing resource
func (h *GenericHandler[Model, RequestDTO, ResponseDTO]) UpdateHandler(
	c *gin.Context,
	validateFunc func(*gin.Context, *RequestDTO) string,
) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.Response(c, response.BadRequestResponse[string]("Invalid ID format"))
		return
	}

	var req RequestDTO
	if err := validateFunc(c, &req); err != "" {
		response.Response(c, response.BadRequestResponse[Model](err))
		return
	}

	resp, err := h.service.Update(c, id, req)
	if err != nil {
		response.Response(c, response.InternalServerErrorResponse[string](err.Error()))
		return
	}

	response.Response(c, response.UpdateDataSuccessResponse(resp))
}

// DeleteHandler handles deleting a resource by ID
func (h *GenericHandler[Model, RequestDTO, ResponseDTO]) DeleteHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.Response(c, response.BadRequestResponse[string]("Invalid ID format"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		response.Response(c, response.InternalServerErrorResponse[string](err.Error()))
		return
	}

	response.Response(c, response.DeleteDataSuccessResponse[any](nil))
}
