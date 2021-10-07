package service

import (
	"context"

	"server-api/module/user/domain/entity"
	"server-api/module/user/domain/repository"
	repositoryTypes "server-api/module/user/infrastructure/repository/types"
	"server-api/module/user/infrastructure/service/types"
)

// UserQueryService handles business logic in the service layer
type UserQueryService struct {
	repository.UserQueryRepositoryInterface
}

// GetUsers returns the users
func (service *UserQueryService) GetUsers(ctx context.Context) ([]entity.User, error) {
	res, err := service.UserQueryRepositoryInterface.SelectUsers()
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetUserByID returns the user by id
func (service *UserQueryService) GetUserByID(ctx context.Context, data types.GetUser) (entity.User, error) {
	var user repositoryTypes.GetUser

	user.ID = data.ID

	res, err := service.UserQueryRepositoryInterface.SelectUserByID(user)
	if err != nil {
		return res, err
	}

	return res, nil
}
