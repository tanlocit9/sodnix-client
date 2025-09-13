package service

import (
	"errors"
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/common/constants"
	"sodnix/apps/server/src/common/mapper"
	"sodnix/apps/server/src/common/repository"
	"sodnix/apps/server/src/common/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenericService implements the CrudServicePort interface
type GenericService[Model any, RequestDTO any, ResponseDTO any] struct {
	repo   repository.GenericRepository[Model]
	mapper mapper.GenericMapper[Model, RequestDTO, ResponseDTO]
}

// NewGenericService creates a new instance of GenericService
func NewGenericService[Model any, RequestDTO any, ResponseDTO any](
	repo repository.GenericRepository[Model],
	mapper mapper.GenericMapper[Model, RequestDTO, ResponseDTO],
) GenericService[Model, RequestDTO, ResponseDTO] {
	return GenericService[Model, RequestDTO, ResponseDTO]{repo, mapper}
}

// Create creates a new model
func (s *GenericService[Model, RequestDTO, ResponseDTO]) Create(c *gin.Context, input RequestDTO) (*ResponseDTO, error) {
	// Validate the input DTO first
	if err := s.mapper.ValidateMapping(input); err != nil {
		return nil, err
	}

	// Convert DTO to entity
	modelPtr, err := s.mapper.ToEntity(input)
	if err != nil {
		return nil, err
	}

	// Handle auditable fields for creation
	if auditable, ok := any(modelPtr).(common.Auditable); ok {
		userID, exists := c.Get(constants.AUTH_USER_ID_KEY)
		if !exists {
			return nil, errors.New("user ID not found in context")
		}
		auditable.SetCreatedBy(userID.(uuid.UUID))
		auditable.SetUpdatedBy(userID.(uuid.UUID))
	}

	// Save to repository
	if err := s.repo.Save(modelPtr); err != nil {
		return nil, err
	}

	// Convert back to response DTO
	responseDTO, err := s.mapper.ToResponseDTO(modelPtr)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

// GetByID retrieves a model by ID
func (s *GenericService[Model, RequestDTO, ResponseDTO]) GetByID(id uuid.UUID, opts *common.GetOptions) (*ResponseDTO, error) {
	model, err := s.repo.FindByID(id, opts)
	if err != nil {
		return nil, err
	}

	responseDTO, err := s.mapper.ToResponseDTO(model)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

// GetAll retrieves all models
func (s *GenericService[Model, RequestDTO, ResponseDTO]) GetAll(params *common.ListQueryParams) (*response.GenericResponse[[]ResponseDTO], error) {
	models, total, err := s.repo.FindAll(params)
	if err != nil {
		return nil, err
	}

	responseDTOs, err := s.mapper.ToResponseDTOList(models)
	if err != nil {
		return nil, err
	}

	return &response.GenericResponse[[]ResponseDTO]{
		Data: &responseDTOs,
		Paging: &response.Paging{
			Page:  params.Page,
			Limit: params.Limit,
			Total: total,
		},
	}, nil
}

// Update updates an existing model
func (s *GenericService[Model, RequestDTO, ResponseDTO]) Update(c *gin.Context, id uuid.UUID, input RequestDTO) (*ResponseDTO, error) {
	existingModel, err := s.repo.FindByID(id, nil)
	// Validate the input DTO first
	if err := s.mapper.ValidateMapping(input); err != nil {
		return nil, err
	}

	// Update the existing model with the input DTO using the mapper
	if err := s.mapper.UpdateEntity(existingModel, input); err != nil {
		return nil, err
	}

	// Set audit fields for update
	if auditModel, ok := any(existingModel).(common.AuditableUpdateOnly); ok {
		userID, exists := c.Get(constants.AUTH_USER_ID_KEY)
		if !exists {
			return nil, errors.New("user ID not found in context")
		}
		auditModel.SetUpdatedBy(userID.(uuid.UUID))
	}

	// Save the updated model
	if err := s.repo.Update(existingModel); err != nil {
		return nil, err
	}

	// Convert to response DTO
	responseDTO, err := s.mapper.ToResponseDTO(existingModel)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

// UpdatePartial updates an existing model with partial data, ignoring specified fields
func (s *GenericService[Model, RequestDTO, ResponseDTO]) UpdatePartial(c *gin.Context, id uuid.UUID, input RequestDTO, ignoreFields []string) (*ResponseDTO, error) {
	// Validate the input DTO first
	if err := s.mapper.ValidateMapping(input); err != nil {
		return nil, err
	}

	// Find existing model
	existingModel, err := s.repo.FindByID(id, nil)
	if err != nil {
		return nil, err // Model not found
	}

	// Update the existing model with partial data
	if err := s.mapper.UpdateEntityPartial(existingModel, input, ignoreFields); err != nil {
		return nil, err
	}

	// Set audit fields for update
	if auditModel, ok := any(existingModel).(common.AuditableUpdateOnly); ok {
		userID, exists := c.Get(constants.AUTH_USER_ID_KEY)
		if !exists {
			return nil, errors.New("user ID not found in context")
		}
		auditModel.SetUpdatedBy(userID.(uuid.UUID))
	}

	// Save the updated model
	if err := s.repo.Update(existingModel); err != nil {
		return nil, err
	}

	// Convert to response DTO
	responseDTO, err := s.mapper.ToResponseDTO(existingModel)
	if err != nil {
		return nil, err
	}

	return responseDTO, nil
}

// Delete deletes a model by ID
func (s *GenericService[Model, RequestDTO, ResponseDTO]) Delete(id uuid.UUID) error {
	// Optionally, check if the model exists before deleting
	_, err := s.repo.FindByID(id, nil)
	if err != nil {
		return errors.New("model not found for deletion")
	}

	return s.repo.Delete(id)
}
