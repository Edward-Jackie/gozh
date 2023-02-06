package controller

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/application/system/service"
)

type SystemEventController struct {
	web.Controller
}

// CreateEvent 数据埋点
// @Summary 创建埋点记录
// @Tags 埋点记录信息
// @Accept json
// @Produce json
// @Router /api/v1/system/event [post]
// @Param createSystemEvent body command.CreateSystemEvent true "新建埋点记录"
// @Success 200 {object} object{code=int,data=entity.SystemEvent}
func (controller SystemEventController) CreateEvent(context *web.Context) {
	createSystemEvent := &command.CreateSystemEvent{}
	_ = context.ShouldBindJSON(createSystemEvent)
	resp, err := service.NewSystemEventService(context).Create(createSystemEvent)
	controller.Response(context, resp, err)
}

func (controller SystemEventController) VideoPlayList(context *web.Context) {
	listSystemEventPlay := &query.ListSystemEventPlay{}
	_ = context.ShouldBindJSON(listSystemEventPlay)
	resp, err := service.NewSystemEventService(context).VideoPlayList(listSystemEventPlay)
	controller.Response(context, resp, err)
}

// HomeStatics 首页统计
// @Summary 首页统计
// @Tags 首页
// @Accept json
// @Produce json
// @Router /api/v1/system/event/home/statics [get]
// @Success 200 {object} object{code=int,data=aggregate.HomeStatics}
func (controller SystemEventController) HomeStatics(context *web.Context) {
	resp, err := service.NewSystemEventService(context).HomeStatics()
	controller.Response(context, resp, err)
}

// HomeChartWithBar 首页统计柱状图
// @Summary 首页统计柱状图
// @Tags 首页
// @Accept json
// @Produce json
// @Router /api/v1/system/event/home/bar [get]
// @Success 200 {object} object{code=int,data=aggregate.HomeStatics}
func (controller SystemEventController) HomeChartWithBar(context *web.Context) {
	getHomeChartBar := &query.GetHomeChartBar{}
	_ = context.ShouldBind(getHomeChartBar)
	resp, err := service.NewSystemEventService(context).HomeChartWithBar(getHomeChartBar)
	controller.Response(context, resp, err)
}
