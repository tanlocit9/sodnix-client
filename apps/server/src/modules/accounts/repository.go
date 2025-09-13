package accounts

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type accountRepository struct {
	repository.GenericRepository[Account]
}

func NewAccountRepository(db *gorm.DB) accountRepository {
	return accountRepository{
		GenericRepository: repository.NewGenericRepository[Account](db),
	}
}

var _ AccountRepositoryPort = (*accountRepository)(nil)
