package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SettingMenuDao struct {
	context *web.Context
}

func NewSettingMenuDao(context *web.Context) *SettingMenuDao {
	return &SettingMenuDao{context: context}
}
