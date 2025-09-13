package repository

import (
	"fmt"
	"sodnix/apps/server/src/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenericRepository implements the CrudRepositoryPort interface
type GenericRepository[T any] struct{ db *gorm.DB }

// NewGenericRepository creates a new instance of GenericRepository
func NewGenericRepository[T any](db *gorm.DB) GenericRepository[T] {
	return GenericRepository[T]{db}
}

// Save saves a model to the database
func (r *GenericRepository[T]) Save(model *T) error {
	return r.db.Create(model).Error
}

// FindByID retrieves a model by ID from the database
func (r *GenericRepository[T]) FindByID(id uuid.UUID, opts *common.GetOptions) (*T, error) {
	var model T
	query := r.db
	if opts != nil && opts.Preload != "" {
		preloads := strings.Split(opts.Preload, ",")
		for _, preload := range preloads {
			query = query.Preload(strings.TrimSpace(preload))
		}
	}
	err := query.First(&model, id).Error
	return &model, err
}

// FindAll retrieves all models from the database
func (r *GenericRepository[T]) FindAll(params *common.ListQueryParams) ([]T, int64, error) {
	var models []T
	var total int64
	query := r.db.Model(new(T))

	if params.Search != "" {
		// This is a simple search, you may need to customize it
		query = query.Where("name LIKE ?", "%"+params.Search+"%")
	}

	if params.SortField != "" {
		order := params.SortOrder
		if order == "" {
			order = "asc"
		}
		query = query.Order(fmt.Sprintf("%s %s", params.SortField, order))
	}

	if params.Preload != "" {
		preloads := strings.Split(params.Preload, ",")
		for _, preload := range preloads {
			query = query.Preload(strings.TrimSpace(preload))
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	err := query.Find(&models).Error
	return models, total, err
}

// Update updates a model in the database
func (r *GenericRepository[T]) Update(model *T) error {
	return r.db.Save(model).Error
}

// Delete deletes a model by ID from the database
func (r *GenericRepository[T]) Delete(id uuid.UUID) error {
	var model *T
	return r.db.Where("id = ?", id).Delete(model).Error
}
