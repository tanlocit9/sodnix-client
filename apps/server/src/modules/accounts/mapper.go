package accounts

import (
	"sodnix/apps/server/src/common/mapper"
)

type accountMapper struct {
	mapper.GenericMapper[Account, AccountRequestDTO, AccountResponseDTO]
}

func NewMapper() *accountMapper {
	return &accountMapper{
		mapper.NewGenericMapper[Account, AccountRequestDTO, AccountResponseDTO](),
	}
}
