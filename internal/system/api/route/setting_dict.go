package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SettingDict() {
	settingDict := web.Default().Group("/api/v1/setting/dict")
	{
		settingDict.POST("/", func(context *gin.Context) {
			controller.SettingDictController{}.Create(web.NewContext(context))
		})
		settingDict.PUT("/", func(context *gin.Context) {
			controller.SettingDictController{}.Update(web.NewContext(context))
		})
		settingDict.GET("/", func(context *gin.Context) {
			controller.SettingDictController{}.Get(web.NewContext(context))
		})
		settingDict.DELETE("/", func(context *gin.Context) {
			controller.SettingDictController{}.Delete(web.NewContext(context))
		})
		settingDict.POST("/list", func(context *gin.Context) {
			controller.SettingDictController{}.List(web.NewContext(context))
		})
	}
}
