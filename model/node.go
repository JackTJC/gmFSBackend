// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 文件节点表
type Node struct {
	NodeID     uint64    `gorm:"column:node_id;primary_key;AUTO_INCREMENT"`             // 节点id
	NodeType   int       `gorm:"column:node_type;NOT NULL"`                             // 节点类型
	NodeName   string    `gorm:"column:node_name;NOT NULL"`                             // 节点名称
	ObjKey     string    `gorm:"column:obj_key;NOT NULL"`                               // 对象在cos中的key
	Content    string    `gorm:"column:content;NOT NULL"`                               // 文件内容,最大64k
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *Node) TableName() string {
	return "node"
}