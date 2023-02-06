package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SysPermissionDao struct {
	context *web.Context
}

func NewSysPermissionDao(context *web.Context) *SysPermissionDao {
	return &SysPermissionDao{context: context}
}
