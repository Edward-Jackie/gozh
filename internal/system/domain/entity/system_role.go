package entity

// SystemRole 角色
type SystemRole struct {
	// 角色ID
	Id int64 `json:"id,string"`
	// 角色编码
	Code string `json:"code"`
	// 角色名称
	Name string `json:"name"`
	// 权限
	Permissions []int64 `json:"permissions"`
	// 备注
	Comment string `json:"comment"`
}

// Update 更新数据
func (systemRole *SystemRole) Update(options map[string]interface{}) {
	// 角色ID
	if id, ok := options["Id"]; ok {
		systemRole.Id = id.(int64)
	}
	// 角色编码
	if code, ok := options["Code"]; ok {
		systemRole.Code = code.(string)
	}
	// 角色名称
	if name, ok := options["Name"]; ok {
		systemRole.Name = name.(string)
	}
	// 权限
	if permissions, ok := options["Permissions"]; ok {
		systemRole.Permissions = permissions.([]int64)
	}
	// 备注
	if comment, ok := options["Comment"]; ok {
		systemRole.Comment = comment.(string)
	}
}
