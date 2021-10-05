package repository

import (
	"golang-api/module/user/domain/entity"
	repositoryTypes "golang-api/module/user/infrastructure/repository/types"
)

// UserQueryRepositoryInterface holds the methods for the user query repository
type UserQueryRepositoryInterface interface {
	SelectUsers() ([]entity.User, error)
	SelectUserByID(data repositoryTypes.GetUser) (entity.User, error)
}
