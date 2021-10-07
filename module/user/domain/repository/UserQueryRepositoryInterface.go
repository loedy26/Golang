package repository

import (
	"server-api/module/user/domain/entity"
	repositoryTypes "server-api/module/user/infrastructure/repository/types"
)

// UserQueryRepositoryInterface holds the methods for the user query repository
type UserQueryRepositoryInterface interface {
	SelectUsers() ([]entity.User, error)
	SelectUserByID(data repositoryTypes.GetUser) (entity.User, error)
}
