package model

import (
	"gorm.io/gorm"
	"time"
)

// VideoInfoItem 视频详情
type VideoInfoItem struct {
	// ID
	Id int64 `json:"id,string"`
	// 视频信息ID
	VideoId int64 `json:"videoId,string"`
	// 标题
	Title string `json:"title"`
	// etag
	Etag string `json:"etag"`
	// 视频地址
	Url string `json:"url"`
	// 缩略图
	Poster string `json:"poster"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 更新时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
