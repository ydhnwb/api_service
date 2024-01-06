package userrepo

import (
	"database/sql"

	"github.com/ydhnwb/api_service/dto"
	"github.com/ydhnwb/api_service/model/user"
)

type UserRepository interface {
	RegisterUser(data dto.UserRegisterRequest) (*user.User, error)
}

type userRepoDependencies struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepoDependencies{
		db: db,
	}
}

func (dep *userRepoDependencies) RegisterUser(data dto.UserRegisterRequest) (*user.User, error) {
	return nil, nil
}
