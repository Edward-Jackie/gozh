package entity

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

// Update 更新数据
func (settingMenu *SettingMenu) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		settingMenu.Id = id.(int64)
	}
	// 菜单英文名称
	if name, ok := options["Name"]; ok {
		settingMenu.Name = name.(string)
	}
	// 路由地址
	if path, ok := options["Path"]; ok {
		settingMenu.Path = path.(string)
	}
	// 类型 menu-菜单 iframe-Iframe link-外链 button-按钮
	if _type, ok := options["Type"]; ok {
		settingMenu.Type = _type.(string)
	}
	// 显示名称
	if title, ok := options["Title"]; ok {
		settingMenu.Title = title.(string)
	}
	// 图标
	if icon, ok := options["Icon"]; ok {
		settingMenu.Icon = icon.(string)
	}
	// 组件视图
	if component, ok := options["Component"]; ok {
		settingMenu.Component = component.(string)
	}
	// 父级ID
	if parentId, ok := options["ParentId"]; ok {
		settingMenu.ParentId = parentId.(int64)
	}
	// 隐藏菜单
	if hidden, ok := options["Hidden"]; ok {
		settingMenu.Hidden = hidden.(int)
	}
	// 隐藏面包屑
	if hideBreadCrumb, ok := options["HideBreadCrumb"]; ok {
		settingMenu.HideBreadCrumb = hideBreadCrumb.(int)
	}
	// 排序
	if sort, ok := options["Sort"]; ok {
		settingMenu.Sort = sort.(int)
	}
}
