package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/filters"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SystemEvent() {
	//添加数据埋点
	web.Default().POST("/api/v1/system/event", func(context *gin.Context) {
		controller.SystemEventController{}.CreateEvent(web.NewContext(context))
	})

	systemEvent := web.Default().Group("/api/v1/system/event").Use(filters.AccessToken())
	{
		//视频播放明细
		systemEvent.POST("/play/list", func(context *gin.Context) {
			controller.SystemEventController{}.VideoPlayList(web.NewContext(context))
		})
		//首页统计
		systemEvent.GET("/home/statics", func(context *gin.Context) {
			controller.SystemEventController{}.HomeStatics(web.NewContext(context))
		})
		//首页柱状图
		systemEvent.GET("/home/bar", func(context *gin.Context) {
			controller.SystemEventController{}.HomeChartWithBar(web.NewContext(context))
		})
	}
}
