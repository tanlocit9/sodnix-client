package accounts

import (
	"sodnix/apps/server/src/common/service"
)

type accountService struct {
	service.GenericService[Account, AccountRequestDTO, AccountResponseDTO]
}

func NewAccountService(repo accountRepository, mapper *accountMapper) accountService {
	return accountService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ AccountServicePort = (*accountService)(nil)
