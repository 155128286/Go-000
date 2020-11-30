package dao

import (
	"database/sql"
	"fmt"
	pkgErr "github.com/pkg/errors"
	"learn-go/Week02/models"
)

type UserDaoImpl struct {
}

func (UserDaoImpl) GetById(id int) (*models.User, error) {
	if id == 1 {
		return nil, pkgErr.Wrap(sql.ErrNoRows, fmt.Sprintf("failed to get user with id: %v", id))
	}

	return &models.User{id, "name", 18, true}, nil

}

func (UserDaoImpl) Create(user *models.User) error {
	return nil
}
