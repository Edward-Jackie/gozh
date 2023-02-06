package dto

import "gozh/internal/system/domain/entity"

type SystemRolePermission struct {
	Id       int64                   `json:"id,string"`
	Label    string                  `json:"label"`
	Children []*SystemRolePermission `json:"children,omitempty"`
}

func TreeMenu(list []*entity.SettingMenu, parentId int64) []*SystemRolePermission {
	res := make([]*SystemRolePermission, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			menu := &SystemRolePermission{
				Id:       v.Id,
				Label:    v.Title,
				Children: make([]*SystemRolePermission, 0),
			}
			menu.Children = TreeMenu(list, v.Id)
			res = append(res, menu)
		}
	}
	return res
}
