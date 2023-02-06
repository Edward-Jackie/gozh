package dao

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/domain/aggregate"
	"gozh/internal/video/infrastructure/model"
)

type VideoInfoDao struct {
	context *web.Context
}

func NewVideoInfoDao(context *web.Context) *VideoInfoDao {
	return &VideoInfoDao{context: context}
}

// GetPlayCount 获取视频播放总量
func (dao *VideoInfoDao) GetPlayCount(videoIds []int64) ([]*aggregate.VideoPlayCount, error) {
	list := make([]*aggregate.VideoPlayCount, 0)
	err := dao.context.Transaction().Context.Model(&model.VideoInfo{}).
		Joins("left join video_info_item on video_info.id=video_info_item.video_id").
		Joins("left join system_event on video_info_item.id = system_event.relation_id").
		Where("video_info.id in ? and system_event.`event` = ? and video_info_item.deleted_at is null", videoIds, "click_video_play").
		Select("video_info.id,count(system_event.id) count").
		Group("video_info.id").
		Scan(&list).Error
	return list, err
}

// GetItemPlayCount 查看视频明细播放量
func (dao *VideoInfoDao) GetItemPlayCount(videoId int64) ([]*aggregate.VideoItemPlayCount, error) {
	list := make([]*aggregate.VideoItemPlayCount, 0)
	err := dao.context.Transaction().Context.Model(&model.VideoInfoItem{}).
		Joins("left join system_event on system_event.relation_id = video_info_item.id").
		Where("video_info_item.video_id = ?", videoId).
		Select("video_info_item.*,count(system_event.id) play_count").
		Group("video_info_item.id").
		Order("video_info_item.created_at desc").
		Scan(&list).Error
	return list, err
}
