package application

import (
	"context"

	"golang-api/module/user/domain/entity"
	"golang-api/module/user/infrastructure/service/types"
)

// UserQueryServiceInterface holds the implementable method for the user query service
type UserQueryServiceInterface interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, data types.GetUser) (entity.User, error)
}
