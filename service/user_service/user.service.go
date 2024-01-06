package userservice

import (
	"os/user"

	"github.com/ydhnwb/api_service/dto"
	userrepo "github.com/ydhnwb/api_service/repository/user_repo"
)

type UserService interface {
	Register(req dto.UserRegisterRequest) (*user.User, error)
}

type userServiceDependencies struct {
	userRepository userrepo.UserRepository
}

func NewUserService(userRepo userrepo.UserRepository) UserService {
	return &userServiceDependencies{
		userRepository: userRepo,
	}
}

func (deps *userServiceDependencies) Register(req dto.UserRegisterRequest) (*user.User, error) {
	return nil, nil
}
