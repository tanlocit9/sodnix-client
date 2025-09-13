package categories

import (
	"sodnix/apps/server/src/common/service"
)

type categoryService struct {
	service.GenericService[Category, CategoryRequestDTO, CategoryResponseDTO]
}

func NewCategoryService(repo categoryRepository, mapper *categoryMapper) categoryService {
	return categoryService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ CategoryServicePort = (*categoryService)(nil)
