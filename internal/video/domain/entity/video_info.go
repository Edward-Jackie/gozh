package entity

import "time"

// VideoInfo 视频
type VideoInfo struct {
	// ID
	Id int64 `json:"id,string"`
	// 视频名称
	Name string `json:"name"`
	// 备注
	Comment string `json:"comment"`
	// 创建人ID
	CreatedBy int64 `json:"createdBy,string"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt"`
	// 类目ID
	CategoryId int64 `json:"categoryId,string" label:"类目ID"`
}

// Update 更新数据
func (videoInfo *VideoInfo) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		videoInfo.Id = id.(int64)
	}
	// 视频名称
	if name, ok := options["Name"]; ok {
		videoInfo.Name = name.(string)
	}
	// 备注
	if comment, ok := options["Comment"]; ok {
		videoInfo.Comment = comment.(string)
	}
	// 类目ID
	if categoryId, ok := options["CategoryId"]; ok {
		videoInfo.CategoryId = categoryId.(int64)
	}
}
