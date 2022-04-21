package db

import (
	"context"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var Node *nodeDB

type nodeDB struct {
}

func (d *nodeDB) Create(ctx context.Context, node *model.Node) error {
	conn := getDbConn(ctx)
	node.CreateTime = time.Now()
	node.UpdateTime = time.Now()
	return conn.Model(node).Create(node).Error
}

func (d *nodeDB) MGetByNodeId(ctx context.Context, nodeIds []int64) (map[int64]*model.Node, error) {
	conn := getDbConn(ctx)
	var nodeList []model.Node
	if err := conn.Model(&model.Node{}).Where("node_id IN (?)", nodeIds).Find(&nodeList).Error; err != nil {
		return nil, err
	}
	ret := make(map[int64]*model.Node)
	for _, node := range nodeList {
		ret[int64(node.NodeID)] = &node
	}
	return ret, nil
}
