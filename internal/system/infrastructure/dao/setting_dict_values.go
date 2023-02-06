package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SettingDictValuesDao struct {
	context *web.Context
}

func NewSettingDictValuesDao(context *web.Context) *SettingDictValuesDao {
	return &SettingDictValuesDao{context: context}
}
