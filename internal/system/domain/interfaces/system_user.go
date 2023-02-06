package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/system/domain/entity"
)

type SystemUserRepository interface {
	Count(*orm.Conditions) (int64, error)                      // 统计记录数
	Save(*entity.SystemUser) (*entity.SystemUser, error)       // 保存数据
	Find(*orm.Conditions) (int64, []*entity.SystemUser, error) // 列表
	FindOne(*orm.Conditions) (*entity.SystemUser, error)       // 查询单条记录
	FindById(int64) (*entity.SystemUser, error)                // 根据ID获取记录
	Delete(*entity.SystemUser) (*entity.SystemUser, error)     // 删除数据
}

type SystemUserDao interface{}
