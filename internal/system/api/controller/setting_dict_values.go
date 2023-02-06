package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/setting/command"
	"gozh/internal/system/application/setting/query"
	"gozh/internal/system/application/setting/service"
)

type SettingDictValuesController struct {
	web.Controller
}

// Create 新增
// @Summary 创建字典值
// @Tags 字典值
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/values [post]
// @Param createSettingDictValues body command.CreateSettingDictValues true "新建字典值"
// @Success 200 {object} object{code=int,data=entity.SettingDictValues}
func (controller SettingDictValuesController) Create(context *web.Context) {
	createSettingDictValues := &command.CreateSettingDictValues{}
	_ = context.ShouldBindJSON(createSettingDictValues)
	resp, err := service.NewSettingDictValuesService(context).Create(createSettingDictValues)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改字典值
// @Tags 字典值
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/values [put]
// @Param updateSettingDictValues body command.UpdateSettingDictValues true "修改字典值"
// @Success 200 {object} object{code=int,data=entity.SettingDictValues}
func (controller SettingDictValuesController) Update(context *web.Context) {
	updateSettingDictValues := &command.UpdateSettingDictValues{}
	_ = context.ShouldBindJSON(updateSettingDictValues)
	resp, err := service.NewSettingDictValuesService(context).Update(updateSettingDictValues)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 字典值列表
// @Tags 字典值
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/values/list [post]
// @Param listSettingDictValues body query.ListSettingDictValues true "修改字典值"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.SettingDictValues}}
func (controller SettingDictValuesController) List(context *web.Context) {
	listSettingDictValues := &query.ListSettingDictValues{}
	_ = context.ShouldBindJSON(listSettingDictValues)
	resp, err := service.NewSettingDictValuesService(context).List(listSettingDictValues)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取字典值
// @Tags 字典值
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/values [get]
// @Param getSettingDictValues query query.GetSettingDictValues true "获取字典值"
// @Success 200 {object} object{code=int,data=entity.SettingDictValues}
func (controller SettingDictValuesController) Get(context *web.Context) {
	getSettingDictValues := &query.GetSettingDictValues{}
	_ = context.ShouldBind(getSettingDictValues)
	resp, err := service.NewSettingDictValuesService(context).Get(getSettingDictValues)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除字典值
// @Tags 字典值
// @Accept json
// @Produce json
// @Router /api/v1/setting/dict/values [delete]
// @Param deleteSettingDictValues body command.DeleteSettingDictValues true "删除字典值"
// @Success 200 {object} object{code=int,data=entity.SettingDictValues}
func (controller SettingDictValuesController) Delete(context *web.Context) {
	deleteSettingDictValues := &command.DeleteSettingDictValues{}
	_ = context.ShouldBindJSON(deleteSettingDictValues)
	resp, err := service.NewSettingDictValuesService(context).Delete(deleteSettingDictValues)
	controller.Response(context, resp, err)
}
