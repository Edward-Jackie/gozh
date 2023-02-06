package video

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/api/route"
)

func Initialization() {
	web.Default().InitRoute(
		route.VideoInfo,
		route.VideoCategory,
	)
}
