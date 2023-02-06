package entity

import (
	"time"
)

// SystemUser 用户表
type SystemUser struct {
	// ID
	Id int64 `json:"id,string"`
	// 用户名称
	UserName string `json:"userName"`
	// 用户昵称
	NickName string `json:"nickName"`
	// 密码
	Password string `json:"-"`
	// 手机号
	Phone string `json:"phone"`
	// 角色
	RoleIds []string `json:"roleIds"`
	// 注册时间
	RegisterTime *time.Time `json:"registerTime"`
	// 最近登录时间
	LastLoginTime *time.Time `json:"lastLoginTime"`
	// 最近登录IP
	LastLoginIp string `json:"lastLoginIp"`
}

var DefaultPassword = "xf123456"

// Update 更新数据
func (systemUser *SystemUser) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		systemUser.Id = id.(int64)
	}
	// 用户名称
	if userName, ok := options["UserName"]; ok {
		systemUser.UserName = userName.(string)
	}
	// 用户昵称
	if nickName, ok := options["NickName"]; ok {
		systemUser.NickName = nickName.(string)
	}
	// 密码
	if password, ok := options["Password"]; ok {
		systemUser.Password = password.(string)
	}
	// 手机号
	if phone, ok := options["Phone"]; ok {
		systemUser.Phone = phone.(string)
	}
	// 角色
	if roleIds, ok := options["RoleIds"]; ok {
		systemUser.RoleIds = roleIds.([]string)
	}
	// 注册时间
	if registerTime, ok := options["RegisterTime"]; ok {
		systemUser.RegisterTime = registerTime.(*time.Time)
	}
	// 最近登录时间
	if lastLoginTime, ok := options["LastLoginTime"]; ok {
		systemUser.LastLoginTime = lastLoginTime.(*time.Time)
	}
	// 最近登录IP
	if lastLoginIp, ok := options["LastLoginIp"]; ok {
		systemUser.LastLoginIp = lastLoginIp.(string)
	}
}
