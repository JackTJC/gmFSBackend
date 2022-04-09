package db

import (
	"errors"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var (
	ErrUserNotFound = errors.New("user not found in db")
	ErrUserExist    = errors.New("user have exist")
)

var User *userDB

type userDB struct {
}

func (d *userDB) CreateUser(user *model.UserInfo) error {
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if err := gormDB.Model(user).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *userDB) GetUserByName(name string) (*model.UserInfo, error) {
	var ret []*model.UserInfo
	err := gormDB.Where("user_name = ?", name).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, nil
	}
	return ret[0], nil
}
