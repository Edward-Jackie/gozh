package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SettingDictValues() {
	settingDictValues := web.Default().Group("/api/v1/setting/dict/values")
	{
		//新增
		settingDictValues.POST("/", func(context *gin.Context) {
			controller.SettingDictValuesController{}.Create(web.NewContext(context))
		})
		//编辑
		settingDictValues.PUT("/", func(context *gin.Context) {
			controller.SettingDictValuesController{}.Update(web.NewContext(context))
		})
		//查询
		settingDictValues.GET("/", func(context *gin.Context) {
			controller.SettingDictValuesController{}.Get(web.NewContext(context))
		})
		//删除
		settingDictValues.DELETE("/", func(context *gin.Context) {
			controller.SettingDictValuesController{}.Delete(web.NewContext(context))
		})
		//列表
		settingDictValues.POST("/list", func(context *gin.Context) {
			controller.SettingDictValuesController{}.List(web.NewContext(context))
		})
	}
}
