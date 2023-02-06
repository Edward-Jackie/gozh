package service

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"

	"github.com/jinzhu/copier"
)

type SettingDictValuesService struct {
	context *web.Context
}

func NewSettingDictValuesService(context *web.Context) *SettingDictValuesService {
	return &SettingDictValuesService{context: context}
}

// Create 新增
func (service *SettingDictValuesService) Create(createSettingDictValues *command.CreateSettingDictValues) (interface{}, error) {
	if err := service.context.Validate(createSettingDictValues); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDictValues := &entity.SettingDictValues{}
	_ = copier.Copy(settingDictValues, createSettingDictValues)
	settingDictValues, err := facade.SettingDictValuesRepository(service.context).Save(settingDictValues)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDictValues, nil
}

// Update 修改
func (service *SettingDictValuesService) Update(updateSettingDictValues *command.UpdateSettingDictValues) (interface{}, error) {
	if err := service.context.Validate(updateSettingDictValues); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDictValuesRepository := facade.SettingDictValuesRepository(service.context)
	settingDictValues, err := settingDictValuesRepository.FindById(updateSettingDictValues.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	settingDictValues.Update(structs.ToMap(updateSettingDictValues))
	settingDictValues, err = facade.SettingDictValuesRepository(service.context).Save(settingDictValues)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDictValues, nil
}

// List 列表
func (service *SettingDictValuesService) List(listSettingDictValues *query.ListSettingDictValues) (interface{}, error) {
	total, list, err := facade.SettingDictValuesRepository(service.context).Find(&orm.Conditions{
		Equal: map[string]interface{}{
			"dictId": listSettingDictValues.DictId,
		},
		Pagination: &orm.Pagination{
			Page:     listSettingDictValues.Page,
			PageSize: listSettingDictValues.PageSize,
		},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return web.NewPagination(total, list), nil
}

// Get 获取信息
func (service *SettingDictValuesService) Get(getSettingDictValues *query.GetSettingDictValues) (interface{}, error) {
	if err := service.context.Validate(getSettingDictValues); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingDictValues, err := facade.SettingDictValuesRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": getSettingDictValues.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingDictValues, nil
}

// Delete 删除
func (service *SettingDictValuesService) Delete(deleteSettingDictValues *command.DeleteSettingDictValues) (interface{}, error) {
	if deleteSettingDictValues.Id <= 0 && len(deleteSettingDictValues.Ids) <= 0 {
		return nil, web.ThrowError(web.ArgError, "id不能为空")
	}
	result := make([]interface{}, 0)
	if deleteSettingDictValues.Id > 0 {
		settingDictValues, err := facade.SettingDictValuesRepository(service.context).FindById(deleteSettingDictValues.Id)
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		settingDictValues, err = facade.SettingDictValuesRepository(service.context).Delete(settingDictValues)
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		result = append(result, settingDictValues)
	}
	if len(deleteSettingDictValues.Ids) > 0 {
		for _, id := range deleteSettingDictValues.Ids {
			settingDictValues, err := facade.SettingDictValuesRepository(service.context).FindById(int64(strutil.MustInt(id)))
			if err != nil {
				return nil, web.ThrowError(web.InternalServerError, err.Error())
			}
			settingDictValues, err = facade.SettingDictValuesRepository(service.context).Delete(settingDictValues)
			if err != nil {
				return nil, web.ThrowError(web.InternalServerError, err.Error())
			}
			result = append(result, settingDictValues)
		}
	}
	return result, nil
}
