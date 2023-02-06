package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SystemUserRepository(context *web.Context) interfaces.SystemUserRepository {
	return repository.NewSystemUserRepository(context)
}

func SystemUserDao(context *web.Context) interfaces.SystemUserDao {
	return dao.NewSystemUserDao(context)
}
