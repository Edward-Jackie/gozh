package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SystemLoggerRepository(context *web.Context) interfaces.SystemLoggerRepository {
	return repository.NewSystemLoggerRepository(context)
}

func SystemLoggerDao(context *web.Context) interfaces.SystemLoggerDao {
	return dao.NewSystemLoggerDao(context)
}
