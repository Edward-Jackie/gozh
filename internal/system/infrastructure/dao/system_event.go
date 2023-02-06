package dao

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/aggregate"
	"gozh/internal/system/infrastructure/model"
)

type SystemEventDao struct {
	context *web.Context
}

func NewSystemEventDao(context *web.Context) *SystemEventDao {
	return &SystemEventDao{context: context}
}

func (dao *SystemEventDao) GetVideoPlayList(options map[string]interface{}) (int64, []*aggregate.SystemEventVideoPlayList, error) {
	list := make([]*aggregate.SystemEventVideoPlayList, 0)
	count := int64(0)
	query := dao.context.Transaction().Context.Model(&model.SystemEvent{})
	query = query.Joins("left join video_info_item on video_info_item.id = system_event.relation_id")
	query = query.Joins("left join video_info on video_info.id = video_info_item.video_id")
	query = query.Where("`event` = ?", "click_video_play")
	query = query.Select("system_event.*,video_info.name as video_name,video_info_item.title as video_title")
	if videoId, ok := options["VideoId"]; ok && videoId.(int64) > 0 {
		query = query.Where("video_info.id = ?", videoId)
	}
	if videoItemId, ok := options["VideoItemId"]; ok && videoItemId.(int64) > 0 {
		query = query.Where("video_info_item.id = ?", videoItemId)
	}
	if playTime, ok := options["PlayTime"]; ok && len(playTime.([]string)) == 2 {
		query = query.Where("date_format(event_time,'%Y-%m-%d') >= ? and date_format(event_time,'%Y-%m-%d') <= ?", playTime.([]string)[0], playTime.([]string)[1])
	}
	if ip, ok := options["Ip"]; ok && ip.(string) != "" {
		query = query.Where("ip = ? ", ip)
	}
	err := query.Count(&count).Error
	if err != nil {
		return count, list, err
	}
	query = query.Order("event_time desc")
	condition := &orm.Conditions{Pagination: &orm.Pagination{Page: 1, PageSize: 10}}
	if page, ok := options["Page"]; ok && page.(int) > 0 {
		condition.Pagination.Page = page.(int)
	}
	if pageSize, ok := options["PageSize"]; ok && pageSize.(int) > 0 {
		condition.Pagination.PageSize = pageSize.(int)
	}
	query = orm.OffsetLimit(query, condition)
	err = query.Scan(&list).Error
	if err != nil {
		return count, list, err
	}
	return count, list, nil
}

// VisitCount 获取访问量统计
func (dao *SystemEventDao) VisitCount(startDate string, overDate string) int64 {
	var count int64
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "visit_video_page", startDate, overDate).Count(&count).Error
	return count
}

// VisitUserCount 获取访问用户数
func (dao *SystemEventDao) VisitUserCount(startDate string, overDate string) int64 {
	var count int64
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "visit_video_page", startDate, overDate).
		Select("count(distinct ip)").
		Scan(&count).Error
	return count
}

// PlayCount 播放量
func (dao *SystemEventDao) PlayCount(startDate, overDate string) int64 {
	var count int64
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "click_video_play", startDate, overDate).Count(&count).Error
	return count
}

// PlayUserCount 获取播放用户数
func (dao *SystemEventDao) PlayUserCount(startDate string, overDate string) int64 {
	var count int64
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "click_video_play", startDate, overDate).
		Select("count(distinct ip)").
		Scan(&count).Error
	return count
}

// VisitCountWithDay 获取访问量统计按天
func (dao *SystemEventDao) VisitCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay {
	list := make([]*aggregate.HomeStaticsDay, 0)
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "visit_video_page", startDate, overDate).
		Group("date_format(event_time,'%Y-%m-%d')").
		Select("date_format(event_time,'%Y-%m-%d') event_date,count(id) count").
		Scan(&list).Error
	return list
}

// VisitUserCountWithDay 获取访问量统计按天
func (dao *SystemEventDao) VisitUserCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay {
	list := make([]*aggregate.HomeStaticsDay, 0)
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "visit_video_page", startDate, overDate).
		Group("date_format(event_time,'%Y-%m-%d')").
		Select("date_format(event_time,'%Y-%m-%d') event_date,count(distinct ip) count").
		Scan(&list).Error
	return list
}

// PlayCountWithDay 获播放量量统计按天
func (dao *SystemEventDao) PlayCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay {
	list := make([]*aggregate.HomeStaticsDay, 0)
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "click_video_play", startDate, overDate).
		Group("date_format(event_time,'%Y-%m-%d')").
		Select("date_format(event_time,'%Y-%m-%d') event_date,count(id) count").
		Scan(&list).Error
	return list
}

// PlayUserCountWithDay 获取播放用户量统计按天
func (dao *SystemEventDao) PlayUserCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay {
	list := make([]*aggregate.HomeStaticsDay, 0)
	_ = dao.context.Transaction().Context.Model(&model.SystemEvent{}).
		Where("`event` = ? and date_format( event_time, '%Y-%m-%d' ) >= ? and date_format( event_time, '%Y-%m-%d' ) <= ?", "click_video_play", startDate, overDate).
		Group("date_format(event_time,'%Y-%m-%d')").
		Select("date_format(event_time,'%Y-%m-%d') event_date,count(distinct ip) count").
		Scan(&list).Error
	return list
}
