/*
Includes DAO related functions
*/
package dao

import (
	"errors"
	pkgErr "github.com/pkg/errors"
	"learn-go/Week02/models"
)

var err = errors.New("Unsupported DAO")

type UserDao interface {
	Create(u *models.User) error
	GetById(id int) (*models.User, error)
}

func New(dao string) (UserDao, error) {
	switch dao {
	case "mysql":
		return UserDaoImpl{}, nil
	default:
		return nil, pkgErr.Wrap(err, "Unsupported DAO "+dao)
	}
}
