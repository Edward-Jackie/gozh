package command

// CreateSettingMenu 新增菜单管理
type CreateSettingMenu struct {
	// 菜单英文名称
	Name string `json:"name" label:"菜单英文名称" validate:"required"`
	// 父级ID
	ParentId int64 `json:"parentId,string" label:"父级ID"`
}
