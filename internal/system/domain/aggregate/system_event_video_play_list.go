package aggregate

import "time"

type SystemEventVideoPlayList struct {
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
	// 视频名称
	VideoName string `json:"videoName"`
	// 视频标题
	VideoTitle string `json:"videoTitle"`
}
