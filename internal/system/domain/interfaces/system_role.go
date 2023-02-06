package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/system/domain/entity"
)

type SystemRoleRepository interface {
	Count(*orm.Conditions) (int64, error)                      // 统计记录数
	Save(*entity.SystemRole) (*entity.SystemRole, error)       // 保存数据
	Find(*orm.Conditions) (int64, []*entity.SystemRole, error) // 列表
	FindOne(*orm.Conditions) (*entity.SystemRole, error)       // 查询单条记录
	FindById(int64) (*entity.SystemRole, error)                // 根据ID获取记录
	Delete(*entity.SystemRole) (*entity.SystemRole, error)     // 删除数据
}

type SystemRoleDao interface{}
