package repository

import (
	"server-api/module/user/domain/entity"
	repositoryTypes "server-api/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryInterface holds the implementable methods for the academic year command repository
type UserCommandRepositoryInterface interface {
	DeleteUserByID(UserID int) error
	InsertUser(data repositoryTypes.CreateUser) (entity.User, error)
	UpdateUserByID(data repositoryTypes.UpdateUser) (entity.User, error)
}
