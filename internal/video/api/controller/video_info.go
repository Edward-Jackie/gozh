package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/application/video/command"
	"gozh/internal/video/application/video/query"
	"gozh/internal/video/application/video/service"
)

type VideoInfoController struct {
	web.Controller
}

// Create 新增
// @Summary 创建视频
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info [post]
// @Param createVideoInfo body command.CreateVideoInfo true "新建视频信息"
// @Success 200 {object} object{code=int,data=entity.VideoInfo}
func (controller VideoInfoController) Create(context *web.Context) {
	createVideoInfo := &command.CreateVideoInfo{}
	_ = context.ShouldBindJSON(createVideoInfo)
	resp, err := service.NewVideoInfoService(context).Create(createVideoInfo)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改视频
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info [put]
// @Param updateVideoInfo body command.UpdateVideoInfo true "修改视频信息"
// @Success 200 {object} object{code=int,data=entity.VideoInfo}
func (controller VideoInfoController) Update(context *web.Context) {
	updateVideoInfo := &command.UpdateVideoInfo{}
	_ = context.ShouldBindJSON(updateVideoInfo)
	resp, err := service.NewVideoInfoService(context).Update(updateVideoInfo)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 视频列表
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info/list [post]
// @Param listVideoInfo body query.ListVideoInfo true "修改视频信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.VideoInfo}}
func (controller VideoInfoController) List(context *web.Context) {
	listVideoInfo := &query.ListVideoInfo{}
	_ = context.ShouldBindJSON(listVideoInfo)
	resp, err := service.NewVideoInfoService(context).List(listVideoInfo)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取视频
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info [get]
// @Param getVideoInfo query query.GetVideoInfo true "获取视频信息"
// @Success 200 {object} object{code=int,data=entity.VideoInfo}
func (controller VideoInfoController) Get(context *web.Context) {
	getVideoInfo := &query.GetVideoInfo{}
	_ = context.ShouldBind(getVideoInfo)
	resp, err := service.NewVideoInfoService(context).Get(getVideoInfo)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除视频
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info [delete]
// @Param deleteVideoInfo body command.DeleteVideoInfo true "删除视频信息"
// @Success 200 {object} object{code=int,data=entity.VideoInfo}
func (controller VideoInfoController) Delete(context *web.Context) {
	deleteVideoInfo := &command.DeleteVideoInfo{}
	_ = context.ShouldBindJSON(deleteVideoInfo)
	resp, err := service.NewVideoInfoService(context).Delete(deleteVideoInfo)
	controller.Response(context, resp, err)
}

// GetItemPlayCount 获取视频明细播放统计
// @Summary 获取视频明细播放统计
// @Tags 视频信息
// @Accept json
// @Produce json
// @Router /api/v1/video/info/item/play/count [get]
// @Param getItemPlayCount query query.GetItemPlayCount true "删除视频信息"
// @Success 200 {object} object{code=int,data=[]aggregate.VideoItemPlayCount}
func (controller VideoInfoController) GetItemPlayCount(context *web.Context) {
	getItemPlayCount := &query.GetItemPlayCount{}
	_ = context.ShouldBind(&getItemPlayCount)
	resp, err := service.NewVideoInfoService(context).GetItemPlayCount(getItemPlayCount)
	controller.Response(context, resp, err)
}
