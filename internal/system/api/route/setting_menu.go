package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SettingMenu() {
	settingMenu := web.Default().Group("/api/v1/setting/menu")
	{
		//新增
		settingMenu.POST("/", func(context *gin.Context) {
			controller.SettingMenuController{}.Create(web.NewContext(context))
		})
		//编辑
		settingMenu.PUT("/", func(context *gin.Context) {
			controller.SettingMenuController{}.Update(web.NewContext(context))
		})
		//查询
		settingMenu.GET("/", func(context *gin.Context) {
			controller.SettingMenuController{}.Get(web.NewContext(context))
		})
		//列表
		settingMenu.POST("/list", func(context *gin.Context) {
			controller.SettingMenuController{}.List(web.NewContext(context))
		})
		//删除
		settingMenu.DELETE("/", func(context *gin.Context) {
			controller.SettingMenuController{}.Delete(web.NewContext(context))
		})
	}
}
