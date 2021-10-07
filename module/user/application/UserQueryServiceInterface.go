package application

import (
	"context"

	"server-api/module/user/domain/entity"
	"server-api/module/user/infrastructure/service/types"
)

// UserQueryServiceInterface holds the implementable method for the user query service
type UserQueryServiceInterface interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, data types.GetUser) (entity.User, error)
}
