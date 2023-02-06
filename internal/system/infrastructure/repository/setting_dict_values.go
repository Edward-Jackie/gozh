package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SettingDictValuesRepository struct {
	context *web.Context
}

func NewSettingDictValuesRepository(context *web.Context) *SettingDictValuesRepository {
	return &SettingDictValuesRepository{context: context}
}

// Count 统计数量
func (repository *SettingDictValuesRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SettingDictValues{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SettingDictValuesRepository) Save(settingDictValues *entity.SettingDictValues) (*entity.SettingDictValues, error) {
	settingDictValuesModel := &model.SettingDictValues{}
	err := copier.Copy(settingDictValuesModel, settingDictValues)
	if err != nil {
		return settingDictValues, err
	}
	if settingDictValuesModel.Id <= 0 {
		settingDictValuesModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return settingDictValues, err
		}
		settingDictValues.Id = settingDictValuesModel.Id
		err = repository.context.Transaction().Context.Create(settingDictValuesModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(settingDictValuesModel).Updates(settingDictValuesModel).Error
	}
	return settingDictValues, err
}

// Find 获取列表
func (repository *SettingDictValuesRepository) Find(conditions *orm.Conditions) (int64, []*entity.SettingDictValues, error) {
	list := make([]*entity.SettingDictValues, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SettingDictValues{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = query.Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *SettingDictValuesRepository) FindOne(conditions *orm.Conditions) (*entity.SettingDictValues, error) {
	settingDictValues := &entity.SettingDictValues{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SettingDictValues{}, conditions)
	err := query.First(settingDictValues).Error
	return settingDictValues, err
}

// FindById 根据ID查询
func (repository *SettingDictValuesRepository) FindById(id int64) (*entity.SettingDictValues, error) {
	settingDictValuesModel := &model.SettingDictValues{}
	settingDictValues := &entity.SettingDictValues{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.SettingDictValues{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(settingDictValuesModel).Error
	if err != nil {
		return settingDictValues, err
	}
	err = copier.Copy(settingDictValues, settingDictValuesModel)
	return settingDictValues, err
}

// Delete 删除
func (repository *SettingDictValuesRepository) Delete(settingDictValues *entity.SettingDictValues) (*entity.SettingDictValues, error) {
	settingDictValuesModel := &model.SettingDictValues{}
	_ = copier.Copy(settingDictValuesModel, settingDictValues)
	err := repository.context.Transaction().Context.Delete(settingDictValuesModel).Error
	return settingDictValues, err
}
