package db

import (
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var File *fileDB

type fileDB struct {
}

func (d *fileDB) CreateFile(file *model.File) error {
	file.CreateTime = time.Now()
	file.UpdateTime = time.Now()
	return gormDB.Model(file).Create(file).Error
}
