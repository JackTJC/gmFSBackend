package db

import (
	"context"
	"errors"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrFileEixst = errors.New("user already have this file")
)

var SecretKey *secretKeyDB

type secretKeyDB struct {
}

func (d *secretKeyDB) Create(ctx context.Context, sk *model.SecretKey) error {
	sk.CreateTime = time.Now()
	sk.UpdateTime = time.Now()
	conn := getDbConn(ctx)
	if err := conn.Model(sk).Create(sk).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ErrFileEixst
		}
		return err
	}
	return nil
}

func (d *secretKeyDB) GetByUIDAndNodeID(ctx context.Context, uid, fileID uint64) (*model.SecretKey, error) {
	ret := &model.SecretKey{}
	conn := getDbConn(ctx)
	err := conn.Model(ret).Where("user_id = ? AND file_id = ?", uid, fileID).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (d *secretKeyDB) GetByUID(ctx context.Context, uid uint64) ([]*model.SecretKey, error) {
	var ret []*model.SecretKey
	conn := getDbConn(ctx)
	err := conn.Model(ret).Where("user_id = ?", uid).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
