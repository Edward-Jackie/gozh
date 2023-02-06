package command

// UpdateSettingMenu 修改菜单管理
type UpdateSettingMenu struct {
	// ID
	Id int64 `json:"id,string" label:"ID" validate:"required"`
	// 菜单英文名称
	Name string `json:"name" label:"菜单英文名称" validate:"required"`
	// 路由地址
	Path string `json:"path" label:"路由地址"`
	// Meta
	Meta struct {
		Type string `json:"type" label:"类型" validate:"required"`
		// 显示名称
		Title string `json:"title" label:"显示名称" validate:"required"`
		// 图标
		Icon string `json:"icon" label:"图标"`
	} `json:"meta" validate:"required"`
	// 组件视图
	Component string `json:"component" label:"组件视图"`
	// 父级ID
	ParentId int64 `json:"parentId,string" label:"父级ID"`
	// 隐藏菜单
	Hidden int `json:"hidden" label:"隐藏菜单"`
	// 隐藏面包屑
	HideBreadCrumb int `json:"hideBreadCrumb" label:"隐藏面包屑"`
	// 排序
	Sort int `json:"sort" label:"排序"`
}
