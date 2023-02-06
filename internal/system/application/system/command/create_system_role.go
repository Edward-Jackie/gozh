package command

// CreateSystemRole 新增角色
type CreateSystemRole struct {
	// 角色编码
	Code string `json:"code" label:"角色编码"`
	// 角色名称
	Name string `json:"name" label:"角色名称"`
	// 备注
	Comment string `json:"comment" label:"备注"`
}
