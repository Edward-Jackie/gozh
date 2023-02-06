package command

// CreateVideoCategory 新增视频类目
type CreateVideoCategory struct {
	// 名称
	Name string `json:"name" label:"名称" validate:"required"`
	// 父级ID
	ParentId int64 `json:"parentId,string" label:"父级ID"`
}
