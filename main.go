package main

import (
	"encoding/json"
	"fmt"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gin-gonic/gin"
	"gozh/internal/system"
	"gozh/internal/video"
	"time"
)

// @title Swagger API
// @version 1.0
// @description Example Api
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	web.Default().InitConfig().InitJob().Middleware(func(c *gin.Context) {
		c.Next()
		resp, exists := c.Get("response")
		if exists {
			respStr, _ := json.Marshal(resp)
			fmt.Println(string(respStr))
		}
	})
	fmt.Println(Uniqid(""))
	//系统模块
	system.Initialization()
	//视频模块
	video.Initialization()
	//启动服务
	web.Default().Start()
}

func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
