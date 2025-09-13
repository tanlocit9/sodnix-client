package transactions

import (
	"sodnix/apps/server/src/common/mapper"
)

type transactionMapper struct {
	mapper.GenericMapper[Transaction, TransactionRequestDTO, TransactionResponseDTO]
}

func NewMapper() *transactionMapper {
	return &transactionMapper{
		mapper.NewGenericMapper[Transaction, TransactionRequestDTO, TransactionResponseDTO](),
	}
}
