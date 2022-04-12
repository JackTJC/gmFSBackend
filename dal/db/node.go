package db

import (
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var Node *nodeDB

type nodeDB struct {
}

func (d *nodeDB) Create(node *model.Node) error {
	node.CreateTime = time.Now()
	node.UpdateTime = time.Now()
	return gormDB.Model(node).Create(node).Error
}
