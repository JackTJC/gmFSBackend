	
// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 分享文件表
type ShareFile struct {
	ID         uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键id
	From       uint64    `gorm:"column:from;NOT NULL"`
	To         uint64    `gorm:"column:to;NOT NULL"`
	FileID     uint64    `gorm:"column:file_id;NOT NULL"`                               // 文件id
	Key        string    `gorm:"column:key;NOT NULL"`                                   // 解密秘钥
	Status     int       `gorm:"column:status;NOT NULL"`                                // 分享是否被处理,1:未处理,2:已处理
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *ShareFile) TableName() string {
	return "share_file"
}
