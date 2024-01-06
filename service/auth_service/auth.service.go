package authservice

import (
	"os/user"

	authrepo "github.com/ydhnwb/api_service/repository/auth_repo"
)

type AuthService interface {
	Login(email string, password string) (*user.User, error)
}

type authServiceDependencies struct {
	authRepo authrepo.AuthRepository
}

func NewAuthService(authRepo authrepo.AuthRepository) AuthService {
	return &authServiceDependencies{
		authRepo: authRepo,
	}
}

func (deps *authServiceDependencies) Login(email string, password string) (*user.User, error) {
	return nil, nil
}
