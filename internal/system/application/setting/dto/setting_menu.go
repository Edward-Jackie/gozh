package dto

import "gozh/internal/system/domain/entity"

type SettingMenu struct {
	// ID
	Id int64 `json:"id,string"`
	// 菜单英文名称
	Name string `json:"name"`
	// 路由地址
	Path string `json:"path"`
	// Meta
	Meta MenuMeta `json:"meta"`
	// 组件视图
	Component string `json:"component,omitempty"`
	// 别名
	Alias string `json:"alias"`
	// 父级ID
	ParentId int64 `json:"parentId,string"`
	// 隐藏菜单
	Hidden int `json:"hidden"`
	// 隐藏面包屑
	HideBreadCrumb int `json:"hideBreadCrumb"`
	// 子集
	Children []*SettingMenu `json:"children,omitempty"`
}

type MenuMeta struct {
	// 显示名称
	Title string `json:"title"`
	// 图标
	Icon string `json:"icon"`
	// 类型 menu-菜单 iframe-Iframe link-外链 button-按钮
	Type string `json:"type"`
}

func TreeMenu(list []*entity.SettingMenu, parentId int64) []*SettingMenu {
	res := make([]*SettingMenu, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			menu := &SettingMenu{
				Id:   v.Id,
				Name: v.Name,
				Path: v.Path,
				Meta: MenuMeta{
					Title: v.Title,
					Icon:  v.Icon,
					Type:  v.Type,
				},
				Component:      v.Component,
				ParentId:       v.ParentId,
				Hidden:         v.Hidden,
				HideBreadCrumb: v.HideBreadCrumb,
				Children:       make([]*SettingMenu, 0),
			}
			menu.Children = TreeMenu(list, v.Id)
			res = append(res, menu)
		}
	}
	return res
}
