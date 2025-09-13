package users

import (
	"sodnix/apps/server/src/common/service"
)

type userService struct {
	service.GenericService[User, UserRequestDTO, UserResponseDTO]
}

func NewUserService(repo userRepository, mapper *userMapper) userService {
	return userService{
		service.NewGenericService(repo.GenericRepository, mapper),
	}
}

var _ UserServicePort = (*userService)(nil)
