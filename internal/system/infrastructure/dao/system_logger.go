package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SystemLoggerDao struct {
	context *web.Context
}

func NewSystemLoggerDao(context *web.Context) *SystemLoggerDao {
	return &SystemLoggerDao{context: context}
}
