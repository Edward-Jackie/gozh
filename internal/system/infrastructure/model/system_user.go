package model

import "time"

// SystemUser 用户表
type SystemUser struct {
	// ID
	Id int64 `json:"id,string"`
	// 用户名称
	UserName string `json:"userName"`
	// 用户昵称
	NickName string `json:"nickName"`
	// 密码
	Password string `json:"password"`
	// 手机号
	Phone string `json:"phone"`
	// 角色
	RoleIds []string `json:"roleIds" gorm:"serializer:json"`
	// 注册时间
	RegisterTime *time.Time `json:"registerTime"`
	// 最近登录时间
	LastLoginTime *time.Time `json:"lastLoginTime"`
	// 最近登录IP
	LastLoginIp string `json:"lastLoginIp"`
	// 创建者
	CreatedBy int64 `json:"createdBy,string"`
	// 修改者
	UpdatedBy int64 `json:"updatedBy,string"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt *time.Time `json:"deletedAt"`
}
