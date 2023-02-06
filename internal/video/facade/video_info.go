package facade

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/domain/interfaces"
	"gozh/internal/video/infrastructure/dao"
	"gozh/internal/video/infrastructure/repository"
)

func VideoInfoRepository(context *web.Context) interfaces.VideoInfoRepository {
	return repository.NewVideoInfoRepository(context)
}

func VideoInfoDao(context *web.Context) interfaces.VideoInfoDao {
	return dao.NewVideoInfoDao(context)
}
