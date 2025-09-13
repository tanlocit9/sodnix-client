package types

import (
	"sodnix/apps/server/src/common/mapper"
)

type typeGroupMapper struct {
	mapper.GenericMapper[TypeGroup, TypeGroupRequestDTO, TypeGroupResponseDTO]
}

func NewTypeGroupMapper() *typeGroupMapper {
	return &typeGroupMapper{
		mapper.NewGenericMapper[TypeGroup, TypeGroupRequestDTO, TypeGroupResponseDTO](),
	}
}
