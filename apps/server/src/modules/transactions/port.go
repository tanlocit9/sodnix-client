package transactions

import (
	"sodnix/apps/server/src/common/port"
)

type TransactionServicePort interface {
	port.CrudServicePort[Transaction, TransactionRequestDTO, TransactionResponseDTO]
}

type TransactionRepositoryPort interface {
	port.CrudRepositoryPort[Transaction]
}
