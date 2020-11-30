package biz

import (
	"learn-go/Week02/dao"
	"learn-go/Week02/models"
)

type UserBiz struct {
	models.User
	Balance int
}

func GetUserInfo(id int) (*UserBiz, error) {
	userDao, err := dao.New("mysql")
	if err != nil {
		return nil, err
	}
	user, err := userDao.GetById(id)
	if err != nil {
		return nil, err
	}
	// get balance
	balance := 100
	userBiz := UserBiz{*user, balance}
	return &userBiz, nil
}
