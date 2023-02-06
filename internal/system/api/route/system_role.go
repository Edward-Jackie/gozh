package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/filters"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"
)

func SystemRole() {
	systemRole := web.Default().Group("/api/v1/system/role").Use(filters.AccessToken())
	{
		//新增
		systemRole.POST("/", func(context *gin.Context) {
			controller.SystemRoleController{}.Create(web.NewContext(context))
		})
		//编辑
		systemRole.PUT("/", func(context *gin.Context) {
			controller.SystemRoleController{}.Update(web.NewContext(context))
		})
		//查询
		systemRole.GET("/", func(context *gin.Context) {
			controller.SystemRoleController{}.Get(web.NewContext(context))
		})
		//删除
		systemRole.DELETE("/", func(context *gin.Context) {
			controller.SystemRoleController{}.Delete(web.NewContext(context))
		})
		//列表
		systemRole.POST("/list", func(context *gin.Context) {
			controller.SystemRoleController{}.List(web.NewContext(context))
		})
		//所有
		systemRole.GET("/all", func(context *gin.Context) {
			controller.SystemRoleController{}.All(web.NewContext(context))
		})
		//获取菜单权限
		systemRole.GET("/permission", func(context *gin.Context) {
			controller.SystemRoleController{}.GetPermissions(web.NewContext(context))
		})
		//保存菜单权限
		systemRole.POST("/permission", func(context *gin.Context) {
			controller.SystemRoleController{}.SavePermissions(web.NewContext(context))
		})
	}
}
