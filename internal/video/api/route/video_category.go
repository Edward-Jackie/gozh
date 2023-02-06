package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gin-gonic/gin"
	"gozh/internal/video/api/controller"
)

func VideoCategory() {
	videoCategory := web.Default().Group("/api/v1/video/category")
	{
		//新增
		videoCategory.POST("/", func(context *gin.Context) {
			controller.VideoCategoryController{}.Create(web.NewContext(context))
		})
		//编辑
		videoCategory.PUT("/", func(context *gin.Context) {
			controller.VideoCategoryController{}.Update(web.NewContext(context))
		})
		//查询
		videoCategory.GET("/", func(context *gin.Context) {
			controller.VideoCategoryController{}.Get(web.NewContext(context))
		})
		//删除
		videoCategory.DELETE("/", func(context *gin.Context) {
			controller.VideoCategoryController{}.Delete(web.NewContext(context))
		})
		//列表
		videoCategory.POST("/list", func(context *gin.Context) {
			controller.VideoCategoryController{}.List(web.NewContext(context))
		})
	}
}
