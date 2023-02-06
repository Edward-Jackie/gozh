package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/video/application/video/command"
	"gozh/internal/video/application/video/query"
	"gozh/internal/video/application/video/service"
)

type VideoCategoryController struct {
	web.Controller
}

// Create 新增
// @Summary 创建视频类目
// @Tags 视频类目信息
// @Accept json
// @Produce json
// @Router /api/v1/videoCategory [post]
// @Param createVideoCategory body command.CreateVideoCategory true "新建视频类目信息"
// @Success 200 {object} object{code=int,data=entity.VideoCategory}
func (controller VideoCategoryController) Create(context *web.Context) {
	createVideoCategory := &command.CreateVideoCategory{}
	_ = context.ShouldBindJSON(createVideoCategory)
	resp, err := service.NewVideoCategoryService(context).Create(createVideoCategory)
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改视频类目
// @Tags 视频类目信息
// @Accept json
// @Produce json
// @Router /api/v1/videoCategory [put]
// @Param updateVideoCategory body command.UpdateVideoCategory true "修改视频类目信息"
// @Success 200 {object} object{code=int,data=entity.VideoCategory}
func (controller VideoCategoryController) Update(context *web.Context) {
	updateVideoCategory := &command.UpdateVideoCategory{}
	_ = context.ShouldBindJSON(updateVideoCategory)
	resp, err := service.NewVideoCategoryService(context).Update(updateVideoCategory)
	controller.Response(context, resp, err)
}

// List 列表
// @Summary 视频类目列表
// @Tags 视频类目信息
// @Accept json
// @Produce json
// @Router /api/v1/videoCategory/list [post]
// @Param listVideoCategory body query.ListVideoCategory true "修改视频类目信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.VideoCategory}}
func (controller VideoCategoryController) List(context *web.Context) {
	listVideoCategory := &query.ListVideoCategory{}
	_ = context.ShouldBindJSON(listVideoCategory)
	resp, err := service.NewVideoCategoryService(context).List(listVideoCategory)
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取视频类目
// @Tags 视频类目信息
// @Accept json
// @Produce json
// @Router /api/v1/videoCategory [get]
// @Param getVideoCategory query query.GetVideoCategory true "获取视频类目信息"
// @Success 200 {object} object{code=int,data=entity.VideoCategory}
func (controller VideoCategoryController) Get(context *web.Context) {
	getVideoCategory := &query.GetVideoCategory{}
	_ = context.ShouldBind(getVideoCategory)
	resp, err := service.NewVideoCategoryService(context).Get(getVideoCategory)
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除视频类目
// @Tags 视频类目信息
// @Accept json
// @Produce json
// @Router /api/v1/videoCategory [delete]
// @Param deleteVideoCategory body command.DeleteVideoCategory true "删除视频类目信息"
// @Success 200 {object} object{code=int,data=entity.VideoCategory}
func (controller VideoCategoryController) Delete(context *web.Context) {
	deleteVideoCategory := &command.DeleteVideoCategory{}
	_ = context.ShouldBindJSON(deleteVideoCategory)
	resp, err := service.NewVideoCategoryService(context).Delete(deleteVideoCategory)
	controller.Response(context, resp, err)
}
