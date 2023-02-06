package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/video/domain/aggregate"
	"gozh/internal/video/domain/entity"
)

type VideoInfoRepository interface {
	Count(*orm.Conditions) (int64, error)                     // 统计记录数
	Save(*entity.VideoInfo) (*entity.VideoInfo, error)        // 保存数据
	Find(*orm.Conditions) (int64, []*entity.VideoInfo, error) // 列表
	FindOne(*orm.Conditions) (*entity.VideoInfo, error)       // 查询单条记录
	FindById(int64) (*entity.VideoInfo, error)                // 根据ID获取记录
	Delete(*entity.VideoInfo) (*entity.VideoInfo, error)      // 删除数据
}

type VideoInfoDao interface {
	GetPlayCount(videoIds []int64) ([]*aggregate.VideoPlayCount, error)      //获取视频组播放总量
	GetItemPlayCount(videoId int64) ([]*aggregate.VideoItemPlayCount, error) //获取视频明细播放统计
}
