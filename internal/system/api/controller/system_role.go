package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/application/system/service"
)

type SystemRoleController struct {
	web.Controller
}

// Create 新增
// @Summary 创建角色
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole [post]
// @Param createSystemRole body command.CreateSystemRole true "新建角色信息"
// @Success 200 {object} object{code=int,data=entity.SystemRole}
func (controller SystemRoleController) Create(context *web.Context) {
	createSystemRole := &command.CreateSystemRole{}
	_ = context.ShouldBindJSON(createSystemRole)
	resp, err := service.NewSystemRoleService(context).Create(createSystemRole)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改角色
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole [put]
// @Param updateSystemRole body command.UpdateSystemRole true "修改角色信息"
// @Success 200 {object} object{code=int,data=entity.SystemRole}
func (controller SystemRoleController) Update(context *web.Context) {
	updateSystemRole := &command.UpdateSystemRole{}
	_ = context.ShouldBindJSON(updateSystemRole)
	resp, err := service.NewSystemRoleService(context).Update(updateSystemRole)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 角色列表
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole/list [post]
// @Param listSystemRole body query.ListSystemRole true "修改角色信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.SystemRole}}
func (controller SystemRoleController) List(context *web.Context) {
	listSystemRole := &query.ListSystemRole{}
	_ = context.ShouldBindJSON(listSystemRole)
	resp, err := service.NewSystemRoleService(context).List(listSystemRole)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取角色
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole [get]
// @Param getSystemRole query query.GetSystemRole true "获取角色信息"
// @Success 200 {object} object{code=int,data=entity.SystemRole}
func (controller SystemRoleController) Get(context *web.Context) {
	getSystemRole := &query.GetSystemRole{}
	_ = context.ShouldBind(getSystemRole)
	resp, err := service.NewSystemRoleService(context).Get(getSystemRole)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除角色
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole [delete]
// @Param deleteSystemRole body command.DeleteSystemRole true "删除角色信息"
// @Success 200 {object} object{code=int,data=entity.SystemRole}
func (controller SystemRoleController) Delete(context *web.Context) {
	deleteSystemRole := &command.DeleteSystemRole{}
	_ = context.ShouldBindJSON(deleteSystemRole)
	resp, err := service.NewSystemRoleService(context).Delete(deleteSystemRole)
	controller.Response(context, resp, err)
}

// GetPermissions 获取菜单权限
// @Summary 获取菜单权限
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole/permission [get]
// @Success 200 {object} object{code=int,data=[]dto.SystemRolePermission}
func (controller SystemRoleController) GetPermissions(context *web.Context) {
	resp, err := service.NewSystemRoleService(context).GetPermissions()
	controller.Response(context, resp, err)
}

// SavePermissions 保存菜单权限
// @Summary 保存菜单权限
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole/permission [post]
// @Success 200 {object} object{code=int,data=entity.SystemRole}
func (controller SystemRoleController) SavePermissions(context *web.Context) {
	savePermission := &command.SavePermissions{}
	_ = context.ShouldBindJSON(savePermission)
	resp, err := service.NewSystemRoleService(context).SavePermissions(savePermission)
	controller.Response(context, resp, err)
}

// All 获取所有角色
// @Summary 获取所有角色
// @Tags 角色信息
// @Accept json
// @Produce json
// @Router /api/v1/systemRole/all [get]
// @Success 200 {object} object{code=int,data=[]entity.CommonLabelValue}
func (controller SystemRoleController) All(context *web.Context) {
	resp, err := service.NewSystemRoleService(context).All()
	controller.Response(context, resp, err)
}
