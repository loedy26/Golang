package repository

import (
	"errors"
	"fmt"
	"strings"

	"server-api/infrastructures/database/mysql/types"
	apiError "server-api/internal/errors"
	"server-api/module/user/domain/entity"
	repositoryTypes "server-api/module/user/infrastructure/repository/types"
)

// UserCommandRepository handles the user command repository logic
type UserCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// DeleteUserByID removes user by id
func (repository *UserCommandRepository) DeleteUserByID(userID int) error {
	user := &entity.User{
		ID: userID,
	}

	// delete user
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		return err
	}

	return nil
}

// InsertUser creates a new user resource
func (repository *UserCommandRepository) InsertUser(data repositoryTypes.CreateUser) (entity.User, error) {
	user := &entity.User{
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		ContactNumber: data.ContactNumber,
	}

	// insert user
	stmt := fmt.Sprintf("INSERT INTO %s (first_name,last_name,contact_number)"+
		"VALUES (:first_name,:last_name,:contact_number)", user.GetModelName())
	res, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		var errStr string

		if strings.Contains(err.Error(), "Duplicate entry") {
			errStr = apiError.DuplicateRecord
		} else {
			errStr = apiError.DatabaseError
		}

		return *user, errors.New(errStr)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return *user, errors.New(apiError.DatabaseError)
	}

	return *user, nil
}

// UpdateUserByID update resource
func (repository *UserCommandRepository) UpdateUserByID(data repositoryTypes.UpdateUser) (entity.User, error) {
	user := &entity.User{
		ID:            data.ID,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		ContactNumber: data.ContactNumber,
	}

	// update user
	stmt := fmt.Sprintf("UPDATE %s SET first_name=:first_name,last_name=:last_name,contact_number=:contact_number "+
		"WHERE id=:id", user.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, user)
	if err != nil {
		fmt.Println(err)
		return *user, errors.New(apiError.DatabaseError)
	}

	return *user, nil
}
