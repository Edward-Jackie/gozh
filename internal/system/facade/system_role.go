package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/interfaces"
	"gozh/internal/system/infrastructure/dao"
	"gozh/internal/system/infrastructure/repository"
)

func SystemRoleRepository(context *web.Context) interfaces.SystemRoleRepository {
	return repository.NewSystemRoleRepository(context)
}

func SystemRoleDao(context *web.Context) interfaces.SystemRoleDao {
	return dao.NewSystemRoleDao(context)
}
