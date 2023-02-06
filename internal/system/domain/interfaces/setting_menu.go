package interfaces

import "gozh/internal/system/domain/entity"

type SettingMenuRepository interface {
	Count(map[string]interface{}) (int64, error)                       // 统计记录数
	Save(*entity.SettingMenu) (*entity.SettingMenu, error)             // 保存数据
	Find(map[string]interface{}) (int64, []*entity.SettingMenu, error) // 列表
	FindOne(map[string]interface{}) (*entity.SettingMenu, error)       // 查询单条记录
	FindById(int64) (*entity.SettingMenu, error)                       // 根据ID获取记录
	Delete(*entity.SettingMenu) (*entity.SettingMenu, error)           // 删除数据
}

type SettingMenuDao interface{}
