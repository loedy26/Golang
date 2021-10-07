package application

import (
	"context"

	"server-api/module/user/domain/entity"
	"server-api/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable method for the user command service
type UserCommandServiceInterface interface {
	CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error)
	DeleteUserByID(userID int) error
	UpdateUserByID(ctx context.Context, data types.UpdateUser) (entity.User, error)
}
