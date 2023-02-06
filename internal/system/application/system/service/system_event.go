package service

import (
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/timex"
	"github.com/jinzhu/copier"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/domain/aggregate"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"time"
)

type SystemEventService struct {
	context *web.Context
}

func NewSystemEventService(context *web.Context) *SystemEventService {
	return &SystemEventService{context: context}
}

// Create 新增
func (service *SystemEventService) Create(createSystemEvent *command.CreateSystemEvent) (interface{}, error) {
	err := service.context.Validate(createSystemEvent)
	if err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemEvent := &entity.SystemEvent{}
	_ = copier.Copy(systemEvent, createSystemEvent)
	systemEvent.EventTime = tools.Now()
	systemEvent.Ip = service.context.ClientIP()
	systemEvent, err = facade.SystemEventRepository(service.context).Save(systemEvent)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return systemEvent, nil
}

// VideoPlayList 播放明细列表
func (service *SystemEventService) VideoPlayList(listSystemEventPlay *query.ListSystemEventPlay) (interface{}, error) {
	count, list, err := facade.SystemEventDao(service.context).GetVideoPlayList(structs.ToMap(listSystemEventPlay))
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return web.NewPagination(count, list), nil
}

// HomeStatics 首页统计
func (service *SystemEventService) HomeStatics() (*aggregate.HomeStatics, error) {
	homeStatics := &aggregate.HomeStatics{}
	systemEventDao := facade.SystemEventDao(service.context)
	//今日
	today := time.Now().Format("2006-01-02")
	homeStatics.TodayVisitCount = systemEventDao.VisitCount(today, today)
	homeStatics.TodayVisitUserCount = systemEventDao.VisitUserCount(today, today)
	homeStatics.TodayPlayCount = systemEventDao.PlayCount(today, today)
	homeStatics.TodayPlayUserCount = systemEventDao.PlayUserCount(today, today)
	//本月
	startDate := time.Now().Format("2006-01") + "-01"
	homeStatics.MonthVisitCount = systemEventDao.VisitCount(startDate, today)
	homeStatics.MonthVisitUserCount = systemEventDao.VisitUserCount(startDate, today)
	homeStatics.MonthPlayCount = systemEventDao.PlayCount(startDate, today)
	homeStatics.MonthPlayUserCount = systemEventDao.PlayUserCount(startDate, today)
	return homeStatics, nil
}

func (service *SystemEventService) HomeChartWithBar(getHomeChartBar *query.GetHomeChartBar) (interface{}, error) {
	if getHomeChartBar.DurationType == "" {
		getHomeChartBar.DurationType = "week"
	}
	var startDate, overDate string
	dateRange := make([]string, 0)
	switch getHomeChartBar.DurationType {
	case "month":
		startDate = timex.NowAddDay(-30).Format("2006-01-02")
		overDate = time.Now().Format("2006-01-02")
		for i := 1; i <= 30; i++ {
			dateRange = append(dateRange, timex.NowAddDay(-30+i).Format("2006-01-02"))
		}
	case "week":
		startDate = timex.NowAddDay(-7).Format("2006-01-02")
		overDate = time.Now().Format("2006-01-02")
		for i := 1; i <= 7; i++ {
			dateRange = append(dateRange, timex.NowAddDay(-7+i).Format("2006-01-02"))
		}
	}
	homeStaticsDayReply := &aggregate.HomeStaticsDayReply{
		Date:    dateRange,
		Statics: make(map[string]*aggregate.HomeStaticsDayInfo),
	}
	//访问量
	systemEventDao := facade.SystemEventDao(service.context)
	visitCountList := systemEventDao.VisitCountWithDay(startDate, overDate)
	for _, item := range visitCountList {
		if _, ok := homeStaticsDayReply.Statics[item.EventDate]; !ok {
			homeStaticsDayReply.Statics[item.EventDate] = &aggregate.HomeStaticsDayInfo{Date: item.EventDate}
		}
		homeStaticsDayReply.Statics[item.EventDate].VisitCount = item.Count
	}
	//访问用户数
	visitUserCountList := systemEventDao.VisitUserCountWithDay(startDate, overDate)
	for _, item := range visitUserCountList {
		if _, ok := homeStaticsDayReply.Statics[item.EventDate]; !ok {
			homeStaticsDayReply.Statics[item.EventDate] = &aggregate.HomeStaticsDayInfo{Date: item.EventDate}
		}
		homeStaticsDayReply.Statics[item.EventDate].VisitUserCount = item.Count
	}
	//播放量
	playCountList := systemEventDao.PlayCountWithDay(startDate, overDate)
	for _, item := range playCountList {
		if _, ok := homeStaticsDayReply.Statics[item.EventDate]; !ok {
			homeStaticsDayReply.Statics[item.EventDate] = &aggregate.HomeStaticsDayInfo{Date: item.EventDate}
		}
		homeStaticsDayReply.Statics[item.EventDate].PlayCount = item.Count
	}
	//播放用户量
	playUserCountList := systemEventDao.PlayUserCountWithDay(startDate, overDate)
	for _, item := range playUserCountList {
		if _, ok := homeStaticsDayReply.Statics[item.EventDate]; !ok {
			homeStaticsDayReply.Statics[item.EventDate] = &aggregate.HomeStaticsDayInfo{Date: item.EventDate}
		}
		homeStaticsDayReply.Statics[item.EventDate].PlayUserCount = item.Count
	}
	return homeStaticsDayReply, nil
}
