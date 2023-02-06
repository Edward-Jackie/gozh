package model

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
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt *time.Time `json:"deletedAt"`
}
