package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SettingMenuRepository(context *web.Context) interfaces.SettingMenuRepository {
	return repository.NewSettingMenuRepository(context)
}

func SettingMenuDao(context *web.Context) interfaces.SettingMenuDao {
	return dao.NewSettingMenuDao(context)
}
