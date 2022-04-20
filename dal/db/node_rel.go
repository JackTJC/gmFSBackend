package db

import (
	"context"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

type nodeRelDB struct {
}

func (d *nodeRelDB) Create(ctx context.Context, m *model.NodeRel) error {
	conn := getDbConn(ctx)
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return conn.Model(&model.NodeRel{}).Create(m).Error
}
