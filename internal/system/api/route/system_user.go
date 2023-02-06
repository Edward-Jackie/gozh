package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/filters"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SystemUser() {
	systemUser := web.Default().Group("/api/v1/system/user").Use(filters.AccessToken())
	{
		//新增
		systemUser.POST("/", func(context *gin.Context) {
			controller.SystemUserController{}.Create(web.NewContext(context))
		})
		//编辑
		systemUser.PUT("/", func(context *gin.Context) {
			controller.SystemUserController{}.Update(web.NewContext(context))
		})
		//查询
		systemUser.GET("/", func(context *gin.Context) {
			controller.SystemUserController{}.Get(web.NewContext(context))
		})
		//删除
		systemUser.DELETE("/", func(context *gin.Context) {
			controller.SystemUserController{}.Delete(web.NewContext(context))
		})
		//列表
		systemUser.POST("/list", func(context *gin.Context) {
			controller.SystemUserController{}.List(web.NewContext(context))
		})
		//修改密码
		systemUser.PUT("/modify/password", func(context *gin.Context) {
			controller.SystemUserController{}.ModifyPassword(web.NewContext(context))
		})
		//重置密码
		systemUser.PUT("/reset/password", func(context *gin.Context) {
			controller.SystemUserController{}.ResetPassword(web.NewContext(context))
		})
	}
}
