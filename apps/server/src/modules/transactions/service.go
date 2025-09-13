package transactions

import (
	"sodnix/apps/server/src/common/service"
)

type transactionService struct {
	service.GenericService[Transaction, TransactionRequestDTO, TransactionResponseDTO]
}

func NewTransactionService(repo transactionRepository, mapper *transactionMapper) transactionService {
	return transactionService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ TransactionServicePort = (*transactionService)(nil)
