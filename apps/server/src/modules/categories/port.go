package categories

import (
	"sodnix/apps/server/src/common/port"
)

type CategoryServicePort interface {
	port.CrudServicePort[Category, CategoryRequestDTO, CategoryResponseDTO]
}

type CategoryRepositoryPort interface {
	port.CrudRepositoryPort[Category]
}
