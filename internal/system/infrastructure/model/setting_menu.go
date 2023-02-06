package model

// SettingMenu 菜单管理
type SettingMenu struct {
	// ID
	Id int64 `json:"id,string"`
	// 菜单英文名称
	Name string `json:"name"`
	// 路由地址
	Path string `json:"path"`
	// 类型 menu-菜单 iframe-Iframe link-外链 button-按钮
	Type string `json:"type"`
	// 显示名称
	Title string `json:"title"`
	// 图标
	Icon string `json:"icon"`
	// 组件视图
	Component string `json:"component"`
	// 父级ID
	ParentId int64 `json:"parentId,string"`
	// 隐藏菜单
	Hidden int `json:"hidden"`
	// 隐藏面包屑
	HideBreadCrumb int `json:"hideBreadCrumb"`
	// 排序
	Sort int `json:"sort"`
}
