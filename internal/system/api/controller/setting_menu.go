package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/application/setting/service"
)

type SettingMenuController struct {
	web.Controller
}

// Create 新增
// @Summary 创建菜单管理
// @Tags 菜单管理信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/menu [post]
// @Param createSettingMenu body command.CreateSettingMenu true "新建菜单管理信息"
// @Success 200 {object} object{code=int,data=entity.SettingMenu}
func (controller SettingMenuController) Create(context *web.Context) {
	createSettingMenu := &command.CreateSettingMenu{}
	_ = context.ShouldBindJSON(createSettingMenu)
	resp, err := service.NewSettingMenuService(context).Create(createSettingMenu)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改菜单管理
// @Tags 菜单管理信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/menu [put]
// @Param updateSettingMenu body command.UpdateSettingMenu true "修改菜单管理信息"
// @Success 200 {object} object{code=int,data=entity.SettingMenu}
func (controller SettingMenuController) Update(context *web.Context) {
	updateSettingMenu := &command.UpdateSettingMenu{}
	_ = context.ShouldBindJSON(updateSettingMenu)
	resp, err := service.NewSettingMenuService(context).Update(updateSettingMenu)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 菜单管理列表
// @Tags 菜单管理信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/menu/list [post]
// @Param listSettingMenu body query.ListSettingMenu true "修改菜单管理信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.SettingMenu}}
func (controller SettingMenuController) List(context *web.Context) {
	listSettingMenu := &query.ListSettingMenu{}
	_ = context.ShouldBindJSON(listSettingMenu)
	resp, err := service.NewSettingMenuService(context).List(listSettingMenu)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取菜单管理
// @Tags 菜单管理信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/menu [get]
// @Param getSettingMenu query query.GetSettingMenu true "获取菜单管理信息"
// @Success 200 {object} object{code=int,data=entity.SettingMenu}
func (controller SettingMenuController) Get(context *web.Context) {
	getSettingMenu := &query.GetSettingMenu{}
	_ = context.ShouldBind(getSettingMenu)
	resp, err := service.NewSettingMenuService(context).Get(getSettingMenu)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除菜单管理
// @Tags 菜单管理信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/menu [delete]
// @Param deleteSettingMenu body command.DeleteSettingMenu true "删除菜单管理信息"
// @Success 200 {object} object{code=int,data=entity.SettingMenu}
func (controller SettingMenuController) Delete(context *web.Context) {
	deleteSettingMenu := &command.DeleteSettingMenu{}
	_ = context.ShouldBindJSON(deleteSettingMenu)
	resp, err := service.NewSettingMenuService(context).Delete(deleteSettingMenu)
	controller.Response(context, resp, err)
}
