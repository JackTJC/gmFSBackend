package db

import (
	"fmt"

	"github.com/JackTJC/gmFS_backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitDB() {
	var err error
	// 配置了安全组，别的ip无法访问,别尝试了 ^_^
	conf := config.GetInstance()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/gm_fs?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.User, conf.MySQL.Passwd, conf.MySQL.Host, conf.MySQL.Port)
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
