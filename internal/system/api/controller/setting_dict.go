package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/application/setting/service"
)

type SettingDictController struct {
	web.Controller
}

// Create 新增
// @Summary 创建数据字典
// @Tags 数据字典信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict [post]
// @Param createSettingDict body command.CreateSettingDict true "新建数据字典信息"
// @Success 200 {object} object{code=int,data=entity.SettingDict}
func (controller SettingDictController) Create(context *web.Context) {
	createSettingDict := &command.CreateSettingDict{}
	_ = context.ShouldBindJSON(createSettingDict)
	resp, err := service.NewSettingDictService(context).Create(createSettingDict)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改数据字典
// @Tags 数据字典信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict [put]
// @Param updateSettingDict body command.UpdateSettingDict true "修改数据字典信息"
// @Success 200 {object} object{code=int,data=entity.SettingDict}
func (controller SettingDictController) Update(context *web.Context) {
	updateSettingDict := &command.UpdateSettingDict{}
	_ = context.ShouldBindJSON(updateSettingDict)
	resp, err := service.NewSettingDictService(context).Update(updateSettingDict)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 数据字典列表
// @Tags 数据字典信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/list [post]
// @Param listSettingDict body query.ListSettingDict true "修改数据字典信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.SettingDict}}
func (controller SettingDictController) List(context *web.Context) {
	listSettingDict := &query.ListSettingDict{}
	_ = context.ShouldBindJSON(listSettingDict)
	resp, err := service.NewSettingDictService(context).List(listSettingDict)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取数据字典
// @Tags 数据字典信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict [get]
// @Param getSettingDict query query.GetSettingDict true "获取数据字典信息"
// @Success 200 {object} object{code=int,data=entity.SettingDict}
func (controller SettingDictController) Get(context *web.Context) {
	getSettingDict := &query.GetSettingDict{}
	_ = context.ShouldBind(getSettingDict)
	resp, err := service.NewSettingDictService(context).Get(getSettingDict)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除数据字典
// @Tags 数据字典信息
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict [delete]
// @Param deleteSettingDict body command.DeleteSettingDict true "删除数据字典信息"
// @Success 200 {object} object{code=int,data=entity.SettingDict}
func (controller SettingDictController) Delete(context *web.Context) {
	deleteSettingDict := &command.DeleteSettingDict{}
	_ = context.ShouldBindJSON(deleteSettingDict)
	resp, err := service.NewSettingDictService(context).Delete(deleteSettingDict)
	controller.Response(context, resp, err)
}
