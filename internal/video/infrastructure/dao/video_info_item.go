package dao

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/infrastructure/model"
)

type VideoInfoItemDao struct {
	context *web.Context
}

func NewVideoInfoItemDao(context *web.Context) *VideoInfoItemDao {
	return &VideoInfoItemDao{context: context}
}

// BatchDelete 批量删除
func (dao *VideoInfoItemDao) BatchDelete(itemIds []int64, videoId int64) error {
	return dao.context.Transaction().Context.Where("video_id = ? and id not in ?", videoId, itemIds).Delete(&model.VideoInfoItem{}).Error
}
