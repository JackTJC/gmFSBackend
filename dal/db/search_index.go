package db

import (
	"context"
	"time"

	"github.com/JackTJC/gmFS_backend/model"
)

var SearchIndex *searchIdxDB

type searchIdxDB struct {
}

func (d *searchIdxDB) MCreate(ctx context.Context, indexs []*model.SearchIndex) error {
	for _, idx := range indexs {
		idx.CreateTime = time.Now()
		idx.UpdateTime = time.Now()
	}
	conn := getDbConn(ctx)
	return conn.Create(indexs).Error
}

func (d *searchIdxDB) SearchByKwAndFileID(ctx context.Context, kw string, fileIDList []uint64) ([]*model.SearchIndex, error) {
	var ret []*model.SearchIndex
	conn := getDbConn(ctx)
	if err := conn.Model(&model.SearchIndex{}).Where("keyword = ? AND file_id IN (?)", kw, fileIDList).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
