package service

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/structs"
	"github.com/jinzhu/copier"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/dto"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
)

type SettingDictService struct {
	context *web.Context
}

func NewSettingDictService(context *web.Context) *SettingDictService {
	return &SettingDictService{context: context}
}

// Create 新增
func (service *SettingDictService) Create(createSettingDict *command.CreateSettingDict) (interface{}, error) {
	if err := service.context.Validate(createSettingDict); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDict := &entity.SettingDict{}
	_ = copier.Copy(settingDict, createSettingDict)
	settingDict, err := facade.SettingDictRepository(service.context).Save(settingDict)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDict, nil
}

// Update 修改
func (service *SettingDictService) Update(updateSettingDict *command.UpdateSettingDict) (interface{}, error) {
	if err := service.context.Validate(updateSettingDict); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDictRepository := facade.SettingDictRepository(service.context)
	settingDict, err := settingDictRepository.FindById(updateSettingDict.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	settingDict.Update(structs.ToMap(updateSettingDict))
	settingDict, err = facade.SettingDictRepository(service.context).Save(settingDict)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDict, nil
}

// List 列表
func (service *SettingDictService) List(listSettingDict *query.ListSettingDict) (interface{}, error) {
	options := structs.ToMap(listSettingDict)
	_, list, err := facade.SettingDictRepository(service.context).Find(options)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return list, nil
}

// Get 获取信息
func (service *SettingDictService) Get(getSettingDict *query.GetSettingDict) (interface{}, error) {
	if err := service.context.Validate(getSettingDict); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	options := structs.ToMap(getSettingDict)
	settingDict, err := facade.SettingDictRepository(service.context).FindOne(options)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	_, dictValues, err := facade.SettingDictValuesRepository(service.context).Find(&orm.Conditions{Equal: map[string]interface{}{"dictId": settingDict.Id}})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	valueResults := make([]dto.DictSelect, 0)
	for _, item := range dictValues {
		valueResults = append(valueResults, dto.DictSelect{
			Label: item.Value,
			Value: item.Key,
		})
	}
	return valueResults, nil
}

// Delete 删除
func (service *SettingDictService) Delete(deleteSettingDict *command.DeleteSettingDict) (interface{}, error) {
	if err := service.context.Validate(deleteSettingDict); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDict, err := facade.SettingDictRepository(service.context).FindById(deleteSettingDict.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	settingDict, err = facade.SettingDictRepository(service.context).Delete(settingDict)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDict, nil
}
