package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SettingDictRepository struct {
	context *web.Context
}

func NewSettingDictRepository(context *web.Context) *SettingDictRepository {
	return &SettingDictRepository{context: context}
}

// Count 统计数量
func (repository *SettingDictRepository) Count(options map[string]interface{}) (int64, error) {
	var total int64
	query := repository.context.Transaction().Context.Model(&model.SettingDict{})
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SettingDictRepository) Save(settingDict *entity.SettingDict) (*entity.SettingDict, error) {
	settingDictModel := &model.SettingDict{}
	err := copier.Copy(settingDictModel, settingDict)
	if err != nil {
		return settingDict, err
	}
	if settingDictModel.Id <= 0 {
		settingDictModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return settingDict, err
		}
		settingDict.Id = settingDictModel.Id
		err = repository.context.Transaction().Context.Create(settingDictModel).Error
	} else {

		err = repository.context.Transaction().Context.Model(settingDictModel).Updates(settingDictModel).Error
	}
	return settingDict, err
}

// Find 获取列表
func (repository *SettingDictRepository) Find(options map[string]interface{}) (int64, []*entity.SettingDict, error) {
	list := make([]*entity.SettingDict, 0)
	models := make([]*model.SettingDict, 0)
	var total int64
	query := repository.context.Transaction().Context.Model(&model.SettingDict{})
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = query.Find(&models).Error
	if err != nil {
		return total, list, err
	}
	//转换
	for _, item := range models {
		list = append(list, &entity.SettingDict{
			Id:   item.Id,
			Code: item.Code,
			Name: item.Name,
		})
	}
	return total, list, err
}

// FindOne 获取单条记录
func (repository *SettingDictRepository) FindOne(options map[string]interface{}) (*entity.SettingDict, error) {
	settingDict := &model.SettingDict{}
	query := repository.context.Transaction().Context.Model(&model.SettingDict{})
	if code, ok := options["Code"]; ok && code.(string) != "" {
		query = query.Where("code = ?", code)
	}
	err := query.First(settingDict).Error
	return &entity.SettingDict{
		Id:   settingDict.Id,
		Code: settingDict.Code,
		Name: settingDict.Name,
	}, err
}

// FindById 根据ID查询
func (repository *SettingDictRepository) FindById(id int64) (*entity.SettingDict, error) {
	settingDictModel := &model.SettingDict{}
	settingDict := &entity.SettingDict{}
	err := repository.context.Transaction().Context.Model(&model.SettingDict{}).Where("id = ?", id).First(settingDictModel).Error
	if err != nil {
		return settingDict, err
	}
	err = copier.Copy(settingDict, settingDictModel)
	return settingDict, err
}

// Delete 删除
func (repository *SettingDictRepository) Delete(settingDict *entity.SettingDict) (*entity.SettingDict, error) {
	settingDictModel := &model.SettingDict{}
	_ = copier.Copy(settingDictModel, settingDict)
	err := repository.context.Transaction().Context.Delete(settingDictModel).Error
	return settingDict, err
}
