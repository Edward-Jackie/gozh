package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/video/domain/entity"
)

type VideoInfoItemRepository interface {
	Count(*orm.Conditions) (int64, error)                         // 统计记录数
	Save(*entity.VideoInfoItem) (*entity.VideoInfoItem, error)    // 保存数据
	Find(*orm.Conditions) (int64, []*entity.VideoInfoItem, error) // 列表
	FindOne(*orm.Conditions) (*entity.VideoInfoItem, error)       // 查询单条记录
	FindById(int64) (*entity.VideoInfoItem, error)                // 根据ID获取记录
	Delete(*entity.VideoInfoItem) (*entity.VideoInfoItem, error)  // 删除数据
}

type VideoInfoItemDao interface {
	BatchDelete(itemIds []int64, videoId int64) error //批量删除
}
