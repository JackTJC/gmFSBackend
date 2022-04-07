// Code generated by sql2gorm. DO NOT EDIT.
package model


// 用户信息表
type UserInfo struct {
	UserID   uint64 `gorm:"column:user_id;primary_key;AUTO_INCREMENT"` // user id
	UserName string `gorm:"column:user_name;NOT NULL"`                 // 用户名字
	Password string `gorm:"column:password;NOT NULL"`                  // 用户密码
	Email    string `gorm:"column:email;NOT NULL"`                     // 用户邮箱
}

func (m *UserInfo) TableName() string {
	return "user_info"
}

