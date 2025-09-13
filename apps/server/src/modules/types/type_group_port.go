package types

import (
	"sodnix/apps/server/src/common/port"
)

type TypeGroupServicePort interface {
	port.CrudServicePort[TypeGroup, TypeGroupRequestDTO, TypeGroupResponseDTO]
}

type TypeGroupRepositoryPort interface {
	port.CrudRepositoryPort[TypeGroup]
}
