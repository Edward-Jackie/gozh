package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/application/system/service"
)

type SystemUserController struct {
	web.Controller
}

// Create 新增
// @Summary 创建用户表
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user [post]
// @Param createSystemUser body command.CreateSystemUser true "新建用户表信息"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) Create(context *web.Context) {
	createSystemUser := &command.CreateSystemUser{}
	_ = context.ShouldBindJSON(createSystemUser)
	resp, err := service.NewSystemUserService(context).Create(createSystemUser)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改用户表
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user [put]
// @Param updateSystemUser body command.UpdateSystemUser true "修改用户表信息"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) Update(context *web.Context) {
	updateSystemUser := &command.UpdateSystemUser{}
	_ = context.ShouldBindJSON(updateSystemUser)
	resp, err := service.NewSystemUserService(context).Update(updateSystemUser)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 用户表列表
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user/list [post]
// @Param listSystemUser body query.ListSystemUser true "修改用户表信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.SystemUser}}
func (controller SystemUserController) List(context *web.Context) {
	listSystemUser := &query.ListSystemUser{}
	_ = context.ShouldBindJSON(listSystemUser)
	resp, err := service.NewSystemUserService(context).List(listSystemUser)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取用户表
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user [get]
// @Param getSystemUser query query.GetSystemUser true "获取用户表信息"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) Get(context *web.Context) {
	getSystemUser := &query.GetSystemUser{}
	_ = context.ShouldBind(getSystemUser)
	resp, err := service.NewSystemUserService(context).Get(getSystemUser)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除用户表
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user [delete]
// @Param deleteSystemUser body command.DeleteSystemUser true "删除用户表信息"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) Delete(context *web.Context) {
	deleteSystemUser := &command.DeleteSystemUser{}
	_ = context.ShouldBindJSON(deleteSystemUser)
	resp, err := service.NewSystemUserService(context).Delete(deleteSystemUser)
	controller.Response(context, resp, err)
}

// ModifyPassword 修改密码
// @Summary 修改密码
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user/modify/password [put]
// @Param modifyPassword body command.ModifyPassword true "修改密码"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) ModifyPassword(context *web.Context) {
	modifyPassword := &command.ModifyPassword{}
	_ = context.ShouldBindJSON(modifyPassword)
	resp, err := service.NewSystemRoleService(context).ModifyPassword(modifyPassword)
	controller.Response(context, resp, err)
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Tags 用户表信息
// @Accept json
// @Produce json
// @Router /api/v1/system/user/reset/password [put]
// @Param resetPassword body command.ResetPassword true "重置密码"
// @Success 200 {object} object{code=int,data=entity.SystemUser}
func (controller SystemUserController) ResetPassword(context *web.Context) {
	resetPassword := &command.ResetPassword{}
	_ = context.ShouldBindJSON(resetPassword)
	resp, err := service.NewSystemUserService(context).ResetPassword(resetPassword)
	controller.Response(context, resp, err)
}

//func (controller SystemUserController) Import(context *web.Context) {
//	file, _ := context.FormFile("file")
//	f, _ := file.Open()
//	rows, err := excel.Read(f)
//
//}
