package transactions

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type transactionRepository struct {
	repository.GenericRepository[Transaction]
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionRepository {
	return transactionRepository{
		GenericRepository: repository.NewGenericRepository[Transaction](db),
		db:                db,
	}
}

var _ TransactionRepositoryPort = (*transactionRepository)(nil)
