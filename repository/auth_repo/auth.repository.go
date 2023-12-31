package authrepo

import (
	"database/sql"
	"os/user"
)

type AuthRepository interface {
	Login(user user.User) (user.User, error)
}

type authRepositoryDependencies struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepositoryDependencies{
		db: db,
	}
}

func (authRepo *authRepositoryDependencies) Login(user user.User) (user.User, error) {
	return user, nil
}
