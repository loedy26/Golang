package application

import (
	"context"

	"golang-api/module/user/domain/entity"
	"golang-api/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable method for the user command service
type UserCommandServiceInterface interface {
	CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error)
	DeleteUserByID(userID int) error
	UpdateUserByID(ctx context.Context, data types.UpdateUser) (entity.User, error)
}
