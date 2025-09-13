package users

import (
	"sodnix/apps/server/src/common/mapper"
)

type userMapper struct {
	mapper.GenericMapper[User, UserRequestDTO, UserResponseDTO]
}

func NewMapper() *userMapper {
	return &userMapper{
		mapper.NewGenericMapper[User, UserRequestDTO, UserResponseDTO](),
	}
}
