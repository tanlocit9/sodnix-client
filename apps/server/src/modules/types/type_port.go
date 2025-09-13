package types

import (
	"sodnix/apps/server/src/common/port"
)

type TypeServicePort interface {
	port.CrudServicePort[Type, TypeRequestDTO, TypeResponseDTO]
}

type TypeRepositoryPort interface {
	port.CrudRepositoryPort[Type]
}
