package types

import (
	"sodnix/apps/server/src/common/service"
)

type typeService struct {
	service.GenericService[Type, TypeRequestDTO, TypeResponseDTO]
}

func NewTypeService(repository typeRepository, mapper *typeMapper) typeService {
	return typeService{
		GenericService: service.NewGenericService(repository.GenericRepository, mapper),
	}
}

// Public
var _ TypeServicePort = (*typeService)(nil)
