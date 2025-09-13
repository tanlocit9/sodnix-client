package types

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type typeGroupRepository struct {
	repository.GenericRepository[TypeGroup]
	// custom methods
}

func NewTypeGroupRepository(db *gorm.DB) *typeGroupRepository {
	return &typeGroupRepository{
		GenericRepository: repository.NewGenericRepository[TypeGroup](db),
	}
}

var _ TypeGroupRepositoryPort = (*typeGroupRepository)(nil)
