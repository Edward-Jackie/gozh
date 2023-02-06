package command

// UpdateSystemRole 修改角色
type UpdateSystemRole struct {
	// 角色ID
	Id int64 `json:"id,string" label:"角色ID"`
	// 角色编码
	Code string `json:"code" label:"角色编码"`
	// 角色名称
	Name string `json:"name" label:"角色名称"`
	// 备注
	Comment string `json:"comment" label:"备注"`
}
