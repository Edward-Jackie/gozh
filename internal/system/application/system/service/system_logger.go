package service

import (
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"gozh/internal/system/infrastructure/cache"
)

type SystemLoggerService struct {
	context *web.Context
}

func NewSystemLoggerService(context *web.Context) *SystemLoggerService {
	return &SystemLoggerService{context: context}
}

// Save 保存日志
func (service *SystemLoggerService) Save(createSystemLogger *command.CreateSystemLogger) error {
	systemLogger := &entity.SystemLogger{
		UserId:     createSystemLogger.UserId,
		Ip:         service.context.ClientIP(),
		Comment:    createSystemLogger.Comment,
		LoggerTime: tools.Now(),
		LoggerType: createSystemLogger.LoggerType,
	}
	if createSystemLogger.UserId > 0 {
		user, _ := cache.GetUserById(service.context, createSystemLogger.UserId)
		systemLogger.NickName = user.NickName
	}
	_, err := facade.SystemLoggerRepository(service.context).Save(systemLogger)
	return err
}
