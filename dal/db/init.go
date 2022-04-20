package db

import (
	"fmt"

	"github.com/JackTJC/gmFS_backend/config"
	"github.com/JackTJC/gmFS_backend/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitDB() {
	var err error
	conf := config.GetInstance()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/gm_fs?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.User, conf.MySQL.Passwd, conf.MySQL.Host, conf.MySQL.Port)
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Sugar.Errorf("open db error:%v", err)
		panic(err)
	}
}

// func Transaction(ctx context.Context, fc func(context.Context) error) {
// 	db := gormDB.WithContext(ctx)
// 	db.Transaction(func(tx *gorm.DB) error {
// 		tx := tx.WithContext(ctx)
// 		db.Transaction(func(tx *gorm.DB) error {
// 			return fc(tx.Statement.Context)
// 		})
// 	})
// }
