package model

import "time"

// SystemLogger
type SystemLogger struct {
	// ID
	Id int64 `json:"id,string"`
	// 用户ID
	UserId int64 `json:"userId,string"`
	// 用户昵称
	NickName string `json:"nickName"`
	// 访问IP
	Ip string `json:"ip"`
	// 日志说明
	Comment string `json:"comment"`
	// 日志时间
	LoggerTime *time.Time `json:"loggerTime"`
	// 类型 1-登录 2-登出 3-用户 4-角色 5-视频
	LoggerType int `json:"loggerType"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt *time.Time `json:"deletedAt"`
}
