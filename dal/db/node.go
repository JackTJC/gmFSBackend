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

func (d *nodeDB) MGetById(nodeIds []int64) (map[int64]*model.Node, error) {
	var nodeList []model.Node
	if err := gormDB.Model(&model.Node{}).Where("node_id IN (?)", nodeIds).Find(&nodeList).Error; err != nil {
		return nil, err
	}
	ret := make(map[int64]*model.Node)
	for _, node := range nodeList {
		ret[int64(node.NodeID)] = &node
	}
	return ret, nil
}
