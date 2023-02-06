package entity

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
}

var (
	LoggerTypeLogin  = 1
	LoggerTypeLogout = 2
	LoggerTypeUser   = 3
	LoggerTypeRole   = 4
)

// Update 更新数据
func (systemLogger *SystemLogger) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		systemLogger.Id = id.(int64)
	}
	// 用户ID
	if userId, ok := options["UserId"]; ok {
		systemLogger.UserId = userId.(int64)
	}
	// 用户昵称
	if nickName, ok := options["NickName"]; ok {
		systemLogger.NickName = nickName.(string)
	}
	// 访问IP
	if ip, ok := options["Ip"]; ok {
		systemLogger.Ip = ip.(string)
	}
	// 日志说明
	if comment, ok := options["Comment"]; ok {
		systemLogger.Comment = comment.(string)
	}
	// 日志时间
	if loggerTime, ok := options["LoggerTime"]; ok {
		systemLogger.LoggerTime = loggerTime.(*time.Time)
	}
	// 类型 1-登录 2-登出 3-用户 4-角色 5-视频
	if loggerType, ok := options["LoggerType"]; ok {
		systemLogger.LoggerType = loggerType.(int)
	}
}
