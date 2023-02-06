package dto

import "time"

type SystemUser struct {
	// ID
	Id int64 `json:"id,string"`
	// 用户名称
	UserName string `json:"userName"`
	// 用户昵称
	NickName string `json:"nickName"`
	// 手机号
	Phone string `json:"phone"`
	// 角色ID
	RoleIds []string `json:"roleIds"`
	// 角色
	Roles string `json:"roles"`
	// 注册时间
	RegisterTime *time.Time `json:"registerTime"`
	// 最近登录时间
	LastLoginTime *time.Time `json:"lastLoginTime"`
	// 最近登录IP
	LastLoginIp string `json:"lastLoginIp"`
}
