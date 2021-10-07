package service

import (
	"context"

	"server-api/module/user/domain/entity"
	"server-api/module/user/domain/repository"
	repositoryTypes "server-api/module/user/infrastructure/repository/types"
	"server-api/module/user/infrastructure/service/types"
)

// UserCommandService handles the user command service logic
type UserCommandService struct {
	repository.UserCommandRepositoryInterface
}

// CreateUser creates a resource and persist it in repository
func (service *UserCommandService) CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error) {
	var user repositoryTypes.CreateUser

	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.ContactNumber = data.ContactNumber

	res, err := service.UserCommandRepositoryInterface.InsertUser(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteUserByID delete user by user id
func (service *UserCommandService) DeleteUserByID(userID int) error {
	err := service.UserCommandRepositoryInterface.DeleteUserByID(userID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserByID updates the resource and persist it in repository
func (service *UserCommandService) UpdateUserByID(ctx context.Context, data types.UpdateUser) (entity.User, error) {
	var user repositoryTypes.UpdateUser

	user.ID = data.ID
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.ContactNumber = data.ContactNumber

	res, err := service.UserCommandRepositoryInterface.UpdateUserByID(user)
	if err != nil {
		return res, err
	}

	return res, nil
}
