package model

import "time"

// VideoCategory 视频类目
type VideoCategory struct {
	// ID
	Id int64 `json:"id,string"`
	// 名称
	Name string `json:"name"`
	// 父级ID
	ParentId int64 `json:"parentId,string"`
	// 层级
	Level int `json:"level"`
	// 路径，json数组格式
	Path string `json:"path"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 编辑时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt *time.Time `json:"deletedAt"`
}
