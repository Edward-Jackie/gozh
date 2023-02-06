package command

type CreateSystemLogger struct {
	// 用户ID
	UserId int64 `json:"userId,string"`
	// 日志说明
	Comment string `json:"comment"`
	// 类型 1-登录 2-登出 3-用户 4-角色 5-视频
	LoggerType int `json:"loggerType"`
}
