package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/video/domain/entity"
)

type VideoCategoryRepository interface {
	Count(*orm.Conditions) (int64, error)                         // 统计记录数
	Save(*entity.VideoCategory) (*entity.VideoCategory, error)    // 保存数据
	Find(*orm.Conditions) (int64, []*entity.VideoCategory, error) // 列表
	FindOne(*orm.Conditions) (*entity.VideoCategory, error)       // 查询单条记录
	FindById(int64) (*entity.VideoCategory, error)                // 根据ID获取记录
	Delete(*entity.VideoCategory) (*entity.VideoCategory, error)  // 删除数据
}

type VideoCategoryDao interface{}
