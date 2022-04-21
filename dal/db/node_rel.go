package db

import (
	"context"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var NodeRel *nodeRelDB

type nodeRelDB struct {
}

func (d *nodeRelDB) Create(ctx context.Context, m *model.NodeRel) error {
	conn := getDbConn(ctx)
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return conn.Model(&model.NodeRel{}).Create(m).Error
}
func (d *nodeRelDB) GetByParent(ctx context.Context, parent int64) ([]*model.NodeRel, error) {
	conn := getDbConn(ctx)
	var ret []*model.NodeRel
	err := conn.Model(&model.NodeRel{}).Where("parent_id = ?", parent).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
