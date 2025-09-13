package users

import (
	"sodnix/apps/server/src/common/port"
)

type UserServicePort interface {
	port.CrudServicePort[User, UserRequestDTO, UserResponseDTO]
}

type UserRepositoryPort interface {
	port.CrudRepositoryPort[User]
}
