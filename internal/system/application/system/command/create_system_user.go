package command

// CreateSystemUser 新增用户表
type CreateSystemUser struct {
	// 用户名称
	UserName string `json:"userName" label:"用户名称"`
	// 用户昵称
	NickName string `json:"nickName" label:"用户昵称"`
	// 密码
	Password string `json:"password" label:"密码"`
	// 手机号
	Phone string `json:"phone" label:"手机号"`
	// 角色
	RoleIds []string `json:"roleIds" label:"角色"`
}
