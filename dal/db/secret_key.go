package db

import (
	"context"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var SecretKey *secretKeyDB

type secretKeyDB struct {
}

func (d *secretKeyDB) Create(ctx context.Context, sk *model.SecretKey) error {
	sk.CreateTime = time.Now()
	sk.UpdateTime = time.Now()
	conn := getDbConn(ctx)
	return conn.Model(sk).Create(sk).Error
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
