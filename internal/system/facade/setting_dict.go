package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SettingDictRepository(context *web.Context) interfaces.SettingDictRepository {
	return repository.NewSettingDictRepository(context)
}

func SettingDictDao(context *web.Context) interfaces.SettingDictDao {
	return dao.NewSettingDictDao(context)
}
