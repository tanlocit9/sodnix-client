package accounts

import (
	"sodnix/apps/server/src/common/port"
)

type AccountServicePort interface {
	port.CrudServicePort[Account, AccountRequestDTO, AccountResponseDTO]
}

type AccountRepositoryPort interface {
	port.CrudRepositoryPort[Account]
}
