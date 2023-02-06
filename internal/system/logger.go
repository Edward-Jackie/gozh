package system

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/service"
)

type Logger struct {
	// 用户ID
	UserId int64 `json:"userId,string"`
	// 日志说明
	Comment string `json:"comment"`
	// 类型 1-登录 2-登出 3-用户 4-角色 5-视频
	LoggerType int `json:"loggerType"`
}

var (
	LoggerTypeVideo = 5
)

func AddLogger(context *web.Context, logger *Logger) error {
	return service.NewSystemLoggerService(context).Save(&command.CreateSystemLogger{
		UserId:     logger.UserId,
		Comment:    logger.Comment,
		LoggerType: logger.LoggerType,
	})
}
