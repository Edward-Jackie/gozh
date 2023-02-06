package command

import "gozh/internal/video/domain/entity"

// UpdateVideoInfo 修改视频
type UpdateVideoInfo struct {
	// ID
	Id int64 `json:"id,string" label:"ID"`
	// 视频名称
	Name string `json:"name" label:"视频名称"`
	// 视频地址，存储为数组json
	Video []*entity.VideoInfoItem `json:"video" label:"视频地址，存储为数组json"`
	// 备注
	Comment string `json:"comment" label:"备注"`
	// 类目ID
	CategoryId int64 `json:"categoryId,string" label:"类目ID"`
}
