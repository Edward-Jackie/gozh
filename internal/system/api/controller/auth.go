package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/service"
)

type AuthController struct {
	web.Controller
}

// Login 登录
// @Summary 登录
// @Tags 用户授权
// @Accept json
// @Produce json
// @Router /api/v1/login [post]
// @Param login body command.Login true "登录信息"
// @Success 200 {object} object{code=int,data=object{token=string,userInfo=entity.SystemUser}}
func (controller AuthController) Login(context *web.Context) {
	login := &command.Login{}
	_ = context.ShouldBindJSON(login)
	resp, err := service.NewAuthService(context).Login(login)
	controller.Response(context, resp, err)
}

// User 获取登录用户信息
// @Summary 获取登录用户信息
// @Tags 用户授权
// @Accept json
// @Produce json
// @Router /api/v1/auth/user [get]
// @Success 200 {object} object{code=int,data=object{token=string,userInfo=entity.SystemUser}}
func (controller AuthController) User(context *web.Context) {
	resp, err := service.NewAuthService(context).User()
	controller.Response(context, resp, err)
}

// Menu 获取登录用户菜单
// @Summary 获取登录用户菜单
// @Tags 用户授权
// @Accept json
// @Produce json
// @Router /api/v1/auth/menu [get]
// @Success 200 {object} object{code=int,data=[]dto.SettingMenu}
func (controller AuthController) Menu(context *web.Context) {
	resp, err := service.NewAuthService(context).Menu()
	controller.Response(context, resp, err)
}

// Dict 获取数据字典
// @Summary 获取登录用户菜单
// @Tags 用户授权
// @Accept json
// @Produce json
// @Router /api/v1/auth/dict [get]
// @Success 200 {object} object{code=int,data=object{}}
func (controller AuthController) Dict(context *web.Context) {
	resp, err := service.NewAuthService(context).Dict()
	controller.Response(context, resp, err)
}
