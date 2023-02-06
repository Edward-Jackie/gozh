package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/domain/interfaces"
	"gozh/internal/video/infrastructure/dao"
	"gozh/internal/video/infrastructure/repository"
)

func VideoInfoItemRepository(context *web.Context) interfaces.VideoInfoItemRepository {
	return repository.NewVideoInfoItemRepository(context)
}

func VideoInfoItemDao(context *web.Context) interfaces.VideoInfoItemDao {
	return dao.NewVideoInfoItemDao(context)
}
