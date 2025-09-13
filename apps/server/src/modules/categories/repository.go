package categories

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type categoryRepository struct {
	repository.GenericRepository[Category]
}

func NewCategoryRepository(db *gorm.DB) categoryRepository {
	return categoryRepository{
		GenericRepository: repository.NewGenericRepository[Category](db),
	}
}

var _ CategoryRepositoryPort = (*categoryRepository)(nil)
