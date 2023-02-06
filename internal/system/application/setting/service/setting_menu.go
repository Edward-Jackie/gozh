package service

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/structs"
	"github.com/jinzhu/copier"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/dto"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"gozh/internal/system/infrastructure/cache"
)

type SettingMenuService struct {
	context *web.Context
}

func NewSettingMenuService(context *web.Context) *SettingMenuService {
	return &SettingMenuService{context: context}
}

// Create 新增
func (service *SettingMenuService) Create(createSettingMenu *command.CreateSettingMenu) (interface{}, error) {
	if err := service.context.Validate(createSettingMenu); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingMenu := &entity.SettingMenu{}
	_ = copier.Copy(settingMenu, createSettingMenu)
	settingMenu, err := facade.SettingMenuRepository(service.context).Save(settingMenu)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	cache.ClearMenu()
	return settingMenu, nil
}

// Update 修改
func (service *SettingMenuService) Update(updateSettingMenu *command.UpdateSettingMenu) (interface{}, error) {
	if err := service.context.Validate(updateSettingMenu); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	settingMenuRepository := facade.SettingMenuRepository(service.context)
	settingMenu, err := settingMenuRepository.FindById(updateSettingMenu.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	settingMenu.Update(structs.ToMap(updateSettingMenu))
	settingMenu.Type = updateSettingMenu.Meta.Type
	settingMenu.Title = updateSettingMenu.Meta.Title
	settingMenu.Icon = updateSettingMenu.Meta.Icon
	settingMenu, err = facade.SettingMenuRepository(service.context).Save(settingMenu)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	cache.ClearMenu()
	return settingMenu, nil
}

// List 列表
func (service *SettingMenuService) List(listSettingMenu *query.ListSettingMenu) (interface{}, error) {
	list, err := cache.GetMenu(service.context)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return dto.TreeMenu(list, 0), nil
}

// Get 获取信息
func (service *SettingMenuService) Get(getSettingMenu *query.GetSettingMenu) (interface{}, error) {
	if err := service.context.Validate(getSettingMenu); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	options := structs.ToMap(getSettingMenu)
	settingMenu, err := facade.SettingMenuRepository(service.context).FindOne(options)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return settingMenu, nil
}

// Delete 删除
func (service *SettingMenuService) Delete(deleteSettingMenu *command.DeleteSettingMenu) (interface{}, error) {
	if err := service.context.Validate(deleteSettingMenu); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	result := make([]*entity.SettingMenu, 0)
	for _, id := range deleteSettingMenu.Ids {
		settingMenu, err := facade.SettingMenuRepository(service.context).FindById(mathutil.MustInt64(id))
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, "查询数据失败")
		}
		//是否有子集
		count, err := facade.SettingMenuRepository(service.context).Count(map[string]interface{}{
			"ParentId": mathutil.MustInt64(id),
		})
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, "查询数据失败")
		}
		if count > 0 {
			return nil, web.ThrowError(web.InternalServerError, settingMenu.Title+"带有子节点，不可删除")
		}
		settingMenu, err = facade.SettingMenuRepository(service.context).Delete(settingMenu)
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		result = append(result, settingMenu)
	}
	cache.ClearMenu()
	return result, nil
}
