package model

import (
	"gorm.io/gorm"
	"time"
)

// SystemRole 角色
type SystemRole struct {
	// 角色ID
	Id int64 `json:"id,string"`
	// 角色名称
	Name string `json:"name"`
	// 备注
	Comment string `json:"comment"`
	// 权限
	Permissions []int64 `json:"permissions" gorm:"serializer:json"`
	// 创建者
	CreatedBy int64 `json:"createdBy,string"`
	// 编辑者
	UpdatedBy int64 `json:"updatedBy,string"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt"`
	// 编辑时间
	UpdatedAt *time.Time `json:"updatedAt"`
	// 删除时间
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
