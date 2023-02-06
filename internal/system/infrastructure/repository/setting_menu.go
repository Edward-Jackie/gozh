package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SettingMenuRepository struct {
	context *web.Context
}

func NewSettingMenuRepository(context *web.Context) *SettingMenuRepository {
	return &SettingMenuRepository{context: context}
}

// Count 统计数量
func (repository *SettingMenuRepository) Count(options map[string]interface{}) (int64, error) {
	var total int64
	query := repository.context.Transaction().Context.Model(&model.SettingMenu{})
	if parentId, ok := options["ParentId"]; ok && parentId.(int64) > 0 {
		query = query.Where("parent_id = ?", parentId)
	}
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SettingMenuRepository) Save(settingMenu *entity.SettingMenu) (*entity.SettingMenu, error) {
	settingMenuModel := &model.SettingMenu{}
	err := copier.Copy(settingMenuModel, settingMenu)
	if err != nil {
		return settingMenu, err
	}
	if settingMenuModel.Id <= 0 {
		settingMenuModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return settingMenu, err
		}
		settingMenu.Id = settingMenuModel.Id
		err = repository.context.Transaction().Context.Create(settingMenuModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(settingMenuModel).Updates(settingMenuModel).Error
	}
	return settingMenu, err
}

// Find 获取列表
func (repository *SettingMenuRepository) Find(options map[string]interface{}) (int64, []*entity.SettingMenu, error) {
	list := make([]*entity.SettingMenu, 0)
	var total int64
	query := repository.context.Transaction().Context.Model(&model.SettingMenu{})
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = query.Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *SettingMenuRepository) FindOne(options map[string]interface{}) (*entity.SettingMenu, error) {
	settingMenu := &entity.SettingMenu{}
	query := repository.context.Transaction().Context.Model(&model.SettingMenu{})
	err := query.First(settingMenu).Error
	return settingMenu, err
}

// FindById 根据ID查询
func (repository *SettingMenuRepository) FindById(id int64) (*entity.SettingMenu, error) {
	settingMenuModel := &model.SettingMenu{}
	settingMenu := &entity.SettingMenu{}
	err := repository.context.Transaction().Context.Where("id=?", id).First(settingMenuModel).Error
	if err != nil {
		return settingMenu, err
	}
	err = copier.Copy(settingMenu, settingMenuModel)
	return settingMenu, err
}

// Delete 删除
func (repository *SettingMenuRepository) Delete(settingMenu *entity.SettingMenu) (*entity.SettingMenu, error) {
	settingMenuModel := &model.SettingMenu{}
	_ = copier.Copy(settingMenuModel, settingMenu)
	err := repository.context.Transaction().Context.Delete(settingMenuModel).Error
	return settingMenu, err
}
