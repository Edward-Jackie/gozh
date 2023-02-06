package dto

type SystemRole struct {
	// 角色ID
	Id int64 `json:"id,string"`
	// 角色编码
	Code string `json:"code"`
	// 角色名称
	Name string `json:"name"`
	// 权限
	Permissions []string `json:"permissions"`
	// 备注
	Comment string `json:"comment"`
}
