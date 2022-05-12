package db

import (
	"context"
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

func (d *userDB) Create(ctx context.Context, user *model.UserInfo) error {
	conn := getDbConn(ctx)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if err := conn.Model(user).Create(user).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ErrUserExist
		}
		return err
	}
	return nil
}

func (d *userDB) GetByName(ctx context.Context, name string) (*model.UserInfo, error) {
	conn := getDbConn(ctx)
	var ret []*model.UserInfo
	err := conn.Where("user_name = ?", name).Find(&ret).Error
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

func (d *userDB) MGetByID(ctx context.Context, uids []uint64) (map[uint64]*model.UserInfo, error) {
	conn := getDbConn(ctx)
	var userList []*model.UserInfo
	if err := conn.Model(&model.UserInfo{}).Where("user_id IN (?)", uids).Find(&userList).Error; err != nil {
		return nil, err
	}
	ret := make(map[uint64]*model.UserInfo)
	for _, user := range userList {
		ret[user.UserID] = user
	}
	return ret, nil
}
