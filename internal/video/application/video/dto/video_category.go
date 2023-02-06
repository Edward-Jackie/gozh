package dto

import "gozh/internal/video/domain/entity"

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
	// 子集
	Children []*VideoCategory `json:"children,omitempty"`
}

func TreeCategory(list []*entity.VideoCategory, parentId int64) []*VideoCategory {
	res := make([]*VideoCategory, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			category := &VideoCategory{
				Id:       v.Id,
				Name:     v.Name,
				Level:    v.Level,
				Path:     v.Path,
				ParentId: v.ParentId,
				Children: make([]*VideoCategory, 0),
			}
			category.Children = TreeCategory(list, v.Id)
			res = append(res, category)
		}
	}
	return res
}
