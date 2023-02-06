package entity

import "fmt"

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
	// 路径，逗号分隔
	Path string `json:"path"`
}

func (videoCategory *VideoCategory) CreatePath(parentPath string, parentId int64) string {
	if parentPath == "" {
		return fmt.Sprintf("%v", parentId)
	} else {
		return parentPath + "," + fmt.Sprintf("%v", parentId)
	}
}

// Update 更新数据
func (videoCategory *VideoCategory) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		videoCategory.Id = id.(int64)
	}
	// 名称
	if name, ok := options["Name"]; ok {
		videoCategory.Name = name.(string)
	}
	// 父级ID
	if parentId, ok := options["ParentId"]; ok {
		videoCategory.ParentId = parentId.(int64)
	}
	// 层级
	if level, ok := options["Level"]; ok {
		videoCategory.Level = level.(int)
	}
	// 路径，json数组格式
	if path, ok := options["Path"]; ok {
		videoCategory.Path = path.(string)
	}
}
