package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SettingDictValuesRepository(context *web.Context) interfaces.SettingDictValuesRepository {
	return repository.NewSettingDictValuesRepository(context)
}

func SettingDictValuesDao(context *web.Context) interfaces.SettingDictValuesDao {
	return dao.NewSettingDictValuesDao(context)
}
