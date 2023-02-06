package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/system/domain/aggregate"
	"gozh/internal/system/domain/entity"
)

type SystemEventRepository interface {
	Count(*orm.Conditions) (int64, error)                       // 统计记录数
	Save(*entity.SystemEvent) (*entity.SystemEvent, error)      // 保存数据
	Find(*orm.Conditions) (int64, []*entity.SystemEvent, error) // 列表
	FindOne(*orm.Conditions) (*entity.SystemEvent, error)       // 查询单条记录
	FindById(int64) (*entity.SystemEvent, error)                // 根据ID获取记录
	Delete(*entity.SystemEvent) (*entity.SystemEvent, error)    // 删除数据
}

type SystemEventDao interface {
	GetVideoPlayList(options map[string]interface{}) (int64, []*aggregate.SystemEventVideoPlayList, error) // 视频播放明细
	VisitCount(startDate string, overDate string) int64                                                    // 获取访问量统计
	VisitUserCount(startDate string, overDate string) int64                                                // 获取访问用户数
	PlayCount(startDate, overDate string) int64                                                            // 播放量统计
	PlayUserCount(startDate string, overDate string) int64                                                 // 获取播放用户数
	VisitCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay
	VisitUserCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay
	PlayCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay
	PlayUserCountWithDay(startDate string, overDate string) []*aggregate.HomeStaticsDay
}
