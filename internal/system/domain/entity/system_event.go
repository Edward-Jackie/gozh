package entity

import "time"

// SystemEvent
type SystemEvent struct {
	// ID
	Id int64 `json:"id"`
	// 埋点事件
	Event string `json:"event"`
	// 访问IP
	Ip string `json:"ip"`
	// ip对应的区域
	Region string `json:"region"`
	// 关联ID
	RelationId int64 `json:"relationId,string"`
	// 时间
	EventTime *time.Time `json:"eventTime"`
}

// Update 更新数据
func (systemEvent *SystemEvent) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		systemEvent.Id = id.(int64)
	}
	// 埋点事件
	if event, ok := options["Event"]; ok {
		systemEvent.Event = event.(string)
	}
	// 访问IP
	if ip, ok := options["Ip"]; ok {
		systemEvent.Ip = ip.(string)
	}
	// ip对应的区域
	if region, ok := options["Region"]; ok {
		systemEvent.Region = region.(string)
	}
	// 关联ID
	if relationId, ok := options["RelationId"]; ok {
		systemEvent.RelationId = relationId.(int64)
	}
	// 时间
	if eventTime, ok := options["EventTime"]; ok {
		systemEvent.EventTime = eventTime.(*time.Time)
	}
}
