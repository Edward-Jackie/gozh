package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/domain/interfaces"
	"gozh/internal/video/infrastructure/dao"
	"gozh/internal/video/infrastructure/repository"
)

func VideoCategoryRepository(context *web.Context) interfaces.VideoCategoryRepository {
	return repository.NewVideoCategoryRepository(context)
}

func VideoCategoryDao(context *web.Context) interfaces.VideoCategoryDao {
	return dao.NewVideoCategoryDao(context)
}
