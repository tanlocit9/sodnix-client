package types

import "sodnix/apps/server/src/common/service"

type typeGroupService struct {
	service.GenericService[TypeGroup, TypeGroupRequestDTO, TypeGroupResponseDTO]
}

func NewTypeGroupService(repository typeGroupRepository, mapper *typeGroupMapper) typeGroupService {
	return typeGroupService{
		GenericService: service.NewGenericService(repository.GenericRepository, mapper),
	}
}

var _ TypeGroupServicePort = (*typeGroupService)(nil)
