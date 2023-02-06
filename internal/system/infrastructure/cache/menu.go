package cache

import (
	"github.com/Edward-Jackie/gotool/pkg/cache"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/repository"
)

var cacheMenuKey = "menu"

var cacheDuration = int64(3600 * 24 * 7)

// GetMenu 获取菜单
func GetMenu(context *web.Context) ([]*entity.SettingMenu, error) {
	data := make([]*entity.SettingMenu, 0)
	_ = cache.Default().Get(cacheMenuKey).ToStruct(&data)
	if len(data) > 0 {
		return data, nil
	}
	_, data, err := repository.NewSettingMenuRepository(context).Find(map[string]interface{}{})
	if err != nil {
		return data, err
	}
	err = cache.Default().Set(cacheMenuKey, data, cacheDuration)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetMenuParentIds(context *web.Context, id int64) ([]int64, error) {
	parentIds := make([]int64, 0)
	menus, err := GetMenu(context)
	if err != nil {
		return parentIds, err
	}
	var menuInfo *entity.SettingMenu
	for _, item := range menus {
		if item.Id == id {
			menuInfo = item
			break
		}
	}
	for _, item := range menus {
		if menuInfo != nil && menuInfo.ParentId == item.Id {
			parentIds = append(parentIds, item.Id)
			//递归父级
			parent, err := GetMenuParentIds(context, item.Id)
			if err == nil && len(parent) > 0 {
				parentIds = append(parentIds, parent...)
			}
		}
	}
	return parentIds, nil
}

// ClearMenu 清除
func ClearMenu() {
	_ = cache.Default().Remove(cacheMenuKey)
}
