package categories

import (
	"sodnix/apps/server/src/common/mapper"
)

type categoryMapper struct {
	mapper.GenericMapper[Category, CategoryRequestDTO, CategoryResponseDTO]
}

func NewMapper() *categoryMapper {
	return &categoryMapper{
		mapper.NewGenericMapper[Category, CategoryRequestDTO, CategoryResponseDTO](),
	}
}
