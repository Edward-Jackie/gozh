package interfaces

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"gozh/internal/system/domain/entity"
)

type SettingDictValuesRepository interface {
	Count(*orm.Conditions) (int64, error)                                // 统计记录数
	Save(*entity.SettingDictValues) (*entity.SettingDictValues, error)   // 保存数据
	Find(*orm.Conditions) (int64, []*entity.SettingDictValues, error)    // 列表
	FindOne(*orm.Conditions) (*entity.SettingDictValues, error)          // 查询单条记录
	FindById(int64) (*entity.SettingDictValues, error)                   // 根据ID获取记录
	Delete(*entity.SettingDictValues) (*entity.SettingDictValues, error) // 删除数据
}

type SettingDictValuesDao interface{}
