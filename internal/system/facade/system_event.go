package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SystemEventRepository(context *web.Context) interfaces.SystemEventRepository {
	return repository.NewSystemEventRepository(context)
}

func SystemEventDao(context *web.Context) interfaces.SystemEventDao {
	return dao.NewSystemEventDao(context)
}
