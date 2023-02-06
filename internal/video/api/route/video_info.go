package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/filters"
	"github.com/gin-gonic/gin"
	"gozh/internal/video/api/controller"
)

func VideoInfo() {
	videoInfo := web.Default().Group("/api/v1/video/info").Use(filters.AccessToken())
	{
		//新增
		videoInfo.POST("/", func(context *gin.Context) {
			controller.VideoInfoController{}.Create(web.NewContext(context))
		})
		//编辑
		videoInfo.PUT("/", func(context *gin.Context) {
			controller.VideoInfoController{}.Update(web.NewContext(context))
		})
		//删除
		videoInfo.DELETE("/", func(context *gin.Context) {
			controller.VideoInfoController{}.Delete(web.NewContext(context))
		})
		//列表
		videoInfo.POST("/list", func(context *gin.Context) {
			controller.VideoInfoController{}.List(web.NewContext(context))
		})
		//获取视频明细播放统计
		videoInfo.GET("/item/play/count", func(context *gin.Context) {
			controller.VideoInfoController{}.GetItemPlayCount(web.NewContext(context))
		})
	}
	//查询视频
	web.Default().GET("/api/v1/video/info", func(context *gin.Context) {
		controller.VideoInfoController{}.Get(web.NewContext(context))
	})
}
