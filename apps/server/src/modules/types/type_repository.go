package types

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type typeRepository struct {
	repository.GenericRepository[Type]
	// custom methods
}

func NewTypeRepository(db *gorm.DB) *typeRepository {
	return &typeRepository{
		GenericRepository: repository.NewGenericRepository[Type](db),
	}
}

var _ TypeRepositoryPort = (*typeRepository)(nil)
