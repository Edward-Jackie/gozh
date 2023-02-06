package aggregate

import "time"

type VideoItemPlayCount struct {
	Id int64 `json:"id,string"`
	// 视频信息ID
	VideoId int64 `json:"videoId,string"`
	// 视频标题
	Title string `json:"title"`
	// etag
	Etag string `json:"etag"`
	// 视频地址
	Url string `json:"url"`
	// 缩略图
	Poster string `json:"poster"`
	// 播放量
	PlayCount int64 `json:"playCount"`
	// 创建日期
	CreatedAt time.Time `json:"createdAt"`
}
