package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SettingDictDao struct {
	context *web.Context
}

func NewSettingDictDao(context *web.Context) *SettingDictDao {
	return &SettingDictDao{context: context}
}
