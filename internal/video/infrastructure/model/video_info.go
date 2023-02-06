package model

import (
	"gorm.io/gorm"
	"time"
)

// VideoInfo 视频
type VideoInfo struct {
	// ID
	Id int64 `json:"id,string"`
	// 视频名称
	Name string `json:"name"`
	// 类目ID
	CategoryId int64 `json:"categoryId,string" label:"类目ID"`
	// 备注
	Comment string `json:"comment"`
	// 创建者
	CreatedBy int64 `json:"createdBy,string"`
	// 编辑者
	UpdatedBy int64 `json:"updatedBy,string"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 编辑时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt gorm.DeletedAt `json:"deletedAt" `
}
