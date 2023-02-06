package command

// UpdateVideoCategory 修改视频类目
type UpdateVideoCategory struct {
	// ID
	Id int64 `json:"id,string" label:"ID"`
	// 名称
	Name string `json:"name" label:"名称"`
	// 父级ID
	ParentId int64 `json:"parentId,string" label:"父级ID"`
	// 层级
	Level int `json:"level" label:"层级"`
	// 路径，json数组格式
	Path string `json:"path" label:"路径，json数组格式"`
}
