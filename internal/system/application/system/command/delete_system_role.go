package command

// DeleteSystemRole 删除角色
type DeleteSystemRole struct {
	// 角色ID
	Id int64 `json:"id,string" label:"角色ID"`
}
