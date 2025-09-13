package types

import (
	"sodnix/apps/server/src/common/mapper"
)

type typeMapper struct {
	mapper.GenericMapper[Type, TypeRequestDTO, TypeResponseDTO]
}

func NewTypeMapper() *typeMapper {
	return &typeMapper{
		mapper.NewGenericMapper[Type, TypeRequestDTO, TypeResponseDTO](),
	}
}
