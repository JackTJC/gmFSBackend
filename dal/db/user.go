package db

import (
	"errors"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ErrUserExist
		}
		return err
	}
	return nil
}

func (d *userDB) GetUserByName(name string) (*model.UserInfo, error) {
	var ret []*model.UserInfo
	err := gormDB.Where("user_name = ?", name).Find(&ret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if len(ret) == 0 {
		return nil, ErrUserNotFound
	}
	return ret[0], nil
}
