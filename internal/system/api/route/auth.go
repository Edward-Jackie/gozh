package route

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/filters"
	"github.com/gin-gonic/gin"
	"gozh/internal/system/api/controller"

	"os"
)

func Auth() {
	//接口文档
	web.Default().GET("/swagger", func(context *gin.Context) {
		mBytes, _ := os.ReadFile("./docs/swagger.json")
		context.Writer.Header().Set("content-type", "application/json;charset=utf8")
		_, _ = context.Writer.Write(mBytes)
	})
	//登录
	web.Default().POST("/api/v1/login", func(context *gin.Context) {
		controller.AuthController{}.Login(web.NewContext(context))
	})
	//获取登录用户信息
	auth := web.Default().Group("/api/v1/auth").Use(filters.AccessToken())
	{
		auth.GET("/user", func(context *gin.Context) {
			controller.AuthController{}.User(web.NewContext(context))
		})
		auth.GET("/menu", func(context *gin.Context) {
			controller.AuthController{}.Menu(web.NewContext(context))
		})
		auth.GET("/dict", func(context *gin.Context) {
			controller.AuthController{}.Dict(web.NewContext(context))
		})
	}
}
