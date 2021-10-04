package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "golang-api/configs/hystrix"
	"golang-api/module/user/domain/entity"
	"golang-api/module/user/domain/repository"
	repositoryTypes "golang-api/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryCircuitBreaker circuit breaker for user command repository
type UserCommandRepositoryCircuitBreaker struct {
	repository.UserCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// DeleteUserByID is the decorator for the the user repository delete by id method
func (repository *UserCommandRepositoryCircuitBreaker) DeleteUserByID(userID int) error {
	hystrix.ConfigureCommand("delete_user_by_id", config.Settings())
	errors := hystrix.Go("delete_user_by_id", func() error {
		err := repository.UserCommandRepositoryInterface.DeleteUserByID(userID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertUser decorator pattern to insert user
func (repository *UserCommandRepositoryCircuitBreaker) InsertUser(data repositoryTypes.CreateUser) (entity.User, error) {
	output := make(chan entity.User, 1)
	hystrix.ConfigureCommand("insert_user", config.Settings())
	errors := hystrix.Go("insert_user", func() error {
		user, err := repository.UserCommandRepositoryInterface.InsertUser(data)
		if err != nil {
			return err
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.User{}, err
	}
}

// UpdateUserByID is the decorator for the user repository update user method
func (repository *UserCommandRepositoryCircuitBreaker) UpdateUserByID(data repositoryTypes.UpdateUser) (entity.User, error) {
	output := make(chan entity.User, 1)
	hystrix.ConfigureCommand("update_user", config.Settings())
	errors := hystrix.Go("update_user", func() error {
		user, err := repository.UserCommandRepositoryInterface.UpdateUserByID(data)
		if err != nil {
			return err
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.User{}, err
	}
}
