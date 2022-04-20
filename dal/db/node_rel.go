package db

import (
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

type nodeRelDB struct {
}

func (d *nodeRelDB) Create(m *model.NodeRel) error {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return gormDB.Model(&model.NodeRel{}).Create(m).Error
}
