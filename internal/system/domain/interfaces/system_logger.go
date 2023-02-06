package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/system/domain/entity"
)

type SystemLoggerRepository interface {
	Count(*orm.Conditions) (int64, error)                        // 统计记录数
	Save(*entity.SystemLogger) (*entity.SystemLogger, error)     // 保存数据
	Find(*orm.Conditions) (int64, []*entity.SystemLogger, error) // 列表
	FindOne(*orm.Conditions) (*entity.SystemLogger, error)       // 查询单条记录
	FindById(int64) (*entity.SystemLogger, error)                // 根据ID获取记录
	Delete(*entity.SystemLogger) (*entity.SystemLogger, error)   // 删除数据
}

type SystemLoggerDao interface{}
